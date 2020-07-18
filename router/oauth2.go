package router

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/traPtitech/traQ/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthParams struct {
	ClientID            string `json:"clientId"`
	CodeChallengeMethod string `json:"codeChallengeMethod"`
	CodeChallenge       string `json:"codeChallenge"`
}

// GetGeneratedCode GET /oauth2/generate/code
func (h *Handlers) GetGeneratedCode(c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		return err
	}

	// セッションIDをすでに持っている場合はCodeChallengeのみ更新
	sessionID, ok := sess.Values["ID"].(string)
	if !ok {
		sess.Options = &h.SessionOption
		sessionID = utils.RandAlphabetAndNumberString(10)
		sess.Values["ID"] = sessionID
	}

	// DBに保存
	codeVerifier := utils.RandAlphabetAndNumberString(43)
	codeVerifierHash := sha256.Sum256([]byte(codeVerifier))
	encoder := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_").WithPadding(base64.NoPadding)

	sess.Values["codeVerifier"] = codeVerifier

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	authParams := &AuthParams{
		ClientID:            h.ClientID,
		CodeChallengeMethod: "S256",
		CodeChallenge:       encoder.EncodeToString(codeVerifierHash[:]),
	}

	return c.JSON(http.StatusCreated, authParams)
}

// Callback GET /oauth2/callback
func (h *Handlers) Callback(code string, c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		return fmt.Errorf("Failed In Getting Session: %w", err)
	}

	interfaceCodeVerifier, ok := sess.Values["codeVerifier"]
	if !ok || interfaceCodeVerifier == nil {
		return errors.New("CodeVerifier IS NULL")
	}
	codeVerifier := interfaceCodeVerifier.(string)

	res, err := o.getAccessToken(code, codeVerifier)
	if err != nil {
		return fmt.Errorf("Failed In Getting AccessToken:%w", err)
	}

	sess.Values["accessToken"] = res.AccessToken
	sess.Values["refreshToken"] = res.RefreshToken

	user, err := o.oauth.GetMe(res.AccessToken)
	if err != nil {
		return fmt.Errorf("Failed In Getting Me: %w", err)
	}

	sess.Values["userID"] = user.Id
	sess.Values["userName"] = user.Name

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("Failed In Save Session: %w", err)
	}

	return nil
}

func (o *OAuth2) getAccessToken(code string, codeVerifier string) (*authResponse, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("client_id", o.clientID)
	form.Set("code", code)
	form.Set("code_verifier", codeVerifier)
	reqBody := strings.NewReader(form.Encode())
	path := o.oauth.BaseURL()
	path.Path += "/oauth2/token"
	req, err := http.NewRequest("POST", path.String(), reqBody)
	if err != nil {
		return &authResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := http.DefaultClient
	res, err := httpClient.Do(req)
	if err != nil {
		return &authResponse{}, err
	}
	if res.StatusCode != 200 {
		return &authResponse{}, fmt.Errorf("Failed In Getting Access Token:(Status:%d %s)", res.StatusCode, res.Status)
	}

	authRes := &authResponse{}
	err = json.NewDecoder(res.Body).Decode(authRes)
	if err != nil {
		return &authResponse{}, fmt.Errorf("Failed In Parsing Json: %w", err)
	}
	return authRes, nil
}
