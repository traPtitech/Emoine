package router

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Emoine/utils"
)

// GetGeneratedCode GET /oauth2/generate/code
func (h *Handlers) GetGeneratedCode(c echo.Context) error {
	sess, err := session.Get("e_session", c)
	if err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}

	sess.Options = &h.SessionOption

	codeVerifier := utils.RandAlphabetAndNumberString(43)
	sess.Values["codeVerifier"] = codeVerifier

	codeVerifierHash := sha256.Sum256([]byte(codeVerifier))
	encoder := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_").WithPadding(base64.NoPadding)

	codeChallengeMethod := "S256"

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}

	return c.Redirect(http.StatusFound, "https://q.trap.jp/api/v3/oauth2/authorize?response_type=code&client_id="+h.ClientID+"&code_challenge="+encoder.EncodeToString(codeVerifierHash[:])+"&code_challenge_method="+codeChallengeMethod)
}

// CallbackHandler GET /api/callback
func (h *Handlers) CallbackHandler(c echo.Context) error {
	sess, err := session.Get("e_session", c)
	if err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}
	if sess.IsNew {
		return echo.ErrUnauthorized
	}

	code := c.QueryParam("code")

	codeVerifier, ok := sess.Values["codeVerifier"].(string)
	if !ok {
		log.Printf("error: type assersion .(string)")
		return echo.ErrInternalServerError
	}

	token, err := requestToken(h.ClientID, code, codeVerifier)
	if err != nil {
		log.Printf("error: %v", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Internal OAuth authentication failed.")
	}

	userMe, err := utils.GetUserMe(token)
	if err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}

	sess.Values["userID"] = userMe.Id.String()
	sess.Values["userName"] = userMe.Name
	sess.Options = &h.SessionOption

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.Redirect(http.StatusFound, "/")
}

func requestToken(clientID, code, codeVerifier string) (token string, err error) {
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("client_id", clientID)
	form.Add("code", code)
	form.Add("code_verifier", codeVerifier)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", "https://q.trap.jp/api/v3/oauth2/token", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode >= 300 {
		return "", err
	}

	data, _ := ioutil.ReadAll(res.Body)
	tokenInfo := TokenResponse{}
	if err := json.Unmarshal(data, &tokenInfo); err != nil {
		return "", err
	}

	return tokenInfo.AccessToken, nil
}
