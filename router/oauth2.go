package router

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/traPtitech/traQ/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type UserID struct {
	Value uuid.UUID `json:"userId"`
}

var verifierCache = cache.New(5*time.Minute, 10*time.Minute)

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

	codeVerifier := utils.RandAlphabetAndNumberString(43)
	codeVerifierHash := sha256.Sum256([]byte(codeVerifier))
	encoder := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_").WithPadding(base64.NoPadding)

	codeChallengeMethod := "S256"

	sess.Values["codeVerifier"] = codeVerifier
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	// セッションIDとをキーにCodeVerifierをキャッシュ
	verifierCache.Set(sessionID, codeVerifier, cache.DefaultExpiration)

	return c.Redirect(http.StatusFound, "https://q.trap.jp/api/v3/oauth2/authorize?response_type=code&client_id="+h.ClientID+"&code_challenge="+encoder.EncodeToString(codeVerifierHash[:])+"&code_challenge_method="+codeChallengeMethod)
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
