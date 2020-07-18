package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/FujishigeTemma/Emoine/utils"
	traQutils "github.com/traPtitech/traQ/utils"

	"github.com/gofrs/uuid"
	jsoniter "github.com/json-iterator/go"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const requestUserStr string = "Request-User"
const authScheme string = "Bearer"

var traQjson = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	TagKey:                 "traq",
}

type OauthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

type UserID struct {
	Value uuid.UUID `json:"userId"`
}

// WatchCallbackMiddleware /callback?code= を監視
func (h *Handlers) WatchCallbackMiddleware() echo.MiddlewareFunc {
	fmt.Println("WatchCallbackMiddleware")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			if path != "/callback" {
				return next(c)
			}
			code := c.QueryParam("code")

			sess, _ := session.Get("session", c)
			sessionID, ok := sess.Values["ID"].(string)
			if !ok {
				return errors.New("sessionID can not parse string")
			}
			codeVerifier, ok := verifierCache.Get(sessionID)
			if !ok {
				return errors.New("codeVerifier is not cached")
			}

			token, err := requestOAuth(h.ClientID, code, codeVerifier.(string))
			if err != nil {
				return err
			}

			// TODO fix
			bytes, _ := utils.GetUserMe(token)
			userID := new(UserID)
			json.Unmarshal(bytes, userID)

			sess.Values["authorization"] = token
			sess.Values["userID"] = userID.Value.String()
			// sess.Options = &h.SessionOption
			err = sess.Save(c.Request(), c.Response())
			if err != nil {
				return err
			}

			return next(c)
		}
	}
}

// TraQUserMiddleware traQユーザーか判定するミドルウェア
func (h *Handlers) TraQUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("TraQUserMiddleware")
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(sess)
		auth, ok := sess.Values["authorization"].(string)
		fmt.Println(auth)
		fmt.Println(ok)
		if !ok {
			sess.Options = &h.SessionOption
			sess.Values["ID"] = traQutils.RandAlphabetAndNumberString(10)
			sess.Save(c.Request(), c.Response())
			return err
		}
		if auth == "" {
			return err
		}
		return next(c)
	}
}

func requestOAuth(clientID, code, codeVerifier string) (token string, err error) {
	fmt.Println("requestOAuth")
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("client_id", clientID)
	form.Add("code", code)
	form.Add("code_verifier", codeVerifier)

	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", "https://q.trap.jp/api/1.0/oauth2/token", body)
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
	oauthRes := new(OauthResponse)
	json.Unmarshal(data, oauthRes)

	token = oauthRes.AccessToken
	return
}
