package router

import (
	"encoding/json"
	"errors"
	"github.com/FujishigeTemma/Emoine/utils"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"os"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

var admins = os.Getenv("ADMINS")

// WatchCallbackMiddleware /callback?code= を監視
func (h *Handlers) WatchCallbackMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if path != "/api/callback" {
			return next(c)
		}
		code := c.QueryParam("code")
		sess, _ := session.Get("e_session", c)
		sessionID, ok := sess.Values["ID"].(string)
		if !ok {
			return errors.New("session_id cannot be parsed as a string")
		}
		codeVerifier, ok := verifierCache.Get(sessionID)
		if !ok {
			return errors.New("code_verifier is not found")
		}

		token, err := requestToken(h.ClientID, code, codeVerifier.(string))
		if err != nil {
			return err
		}

		bytes, _ := utils.GetUserMe(token)
		userID := new(UserID)
		if err := json.Unmarshal(bytes, userID); err != nil {
			return err
		}

		sess.Values["accessToken"] = token
		sess.Values["userID"] = userID.Value.String()
		sess.Options = &h.SessionOption

		sessionCache.Add(userID.Value.String(), token, cache.DefaultExpiration)

		err = sess.Save(c.Request(), c.Response())
		if err != nil {
			return internalServerError(err)
		}

		return next(c)
	}
}

// TraQUserMiddleware traQユーザーか判定するミドルウェア
func (h *Handlers) IsTraQUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("e_session", c)
		if err != nil {
			return unauthorized(err)
		}
		auth, ok := sess.Values["accessToken"].(string)
		if !ok {
			sess.Options = &h.SessionOption
			sess.Values["ID"] = utils.RandAlphabetAndNumberString(10)
			sess.Save(c.Request(), c.Response())
			return unauthorized(err)
		}
		if auth == "" {
			return unauthorized(err)
		}
		setRequestUserIsAdmin(c)
		return next(c)
	}
}

func getRequestUserID(c echo.Context) (uuid.UUID, error) {
	sess, err := session.Get("e_session", c)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromString(sess.Values["userID"].(string))
}

func setRequestUserIsAdmin(c echo.Context) {
	userID, _ := getRequestUserID(c)
	c.Set("IsAdmin", strings.Contains(admins, userID.String()))
}

// AdminUserMiddleware 管理者ユーザーか判定するミドルウェア
func (h *Handlers) AdminUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := getRequestUserIsAdmin(c)

		// 判定
		if !isAdmin {
			return forbidden(
				errors.New("not admin"),
				message("You are not admin user."),
				specification("Only admin user can request."),
			)
		}

		return next(c)
	}
}

func getRequestUserIsAdmin(c echo.Context) bool {
	return c.Get("IsAdmin").(bool)
}

func getRequestUserToken(c echo.Context) (string, error) {
	sess, err := session.Get("e_session", c)
	if err != nil {
		return "", err
	}
	token, ok := sess.Values["accessToken"].(string)
	if !ok {
		return "", errors.New("error")
	}

	return token, nil
}
