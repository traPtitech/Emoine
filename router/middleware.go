package router

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/FujishigeTemma/Emoine/utils"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
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

		userMe, err := utils.GetUserMe(token)
		if err != nil {
			return err
		}

		sess.Values["accessToken"] = token
		sess.Values["userID"] = userMe.Id.String()
		sess.Values["userName"] = userMe.Name
		sess.Options = &h.SessionOption

		sessionCache.Add(userMe.Id.String(), token, cache.DefaultExpiration)

		err = sess.Save(c.Request(), c.Response())
		if err != nil {
			return internalServerError(err)
		}

		return c.Redirect(http.StatusFound, "/")
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

func GetRequestUserID(c echo.Context) (uuid.UUID, error) {
	sess, err := session.Get("e_session", c)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromString(sess.Values["userID"].(string))
}

func GetRequestUserName(c echo.Context) (string, error) {
	sess, err := session.Get("e_session", c)
	if err != nil {
		return "", err
	}
	userName := sess.Values["userName"]
	if userName == nil {
		return "", nil
	}
	return userName.(string), nil
}

func setRequestUserIsAdmin(c echo.Context) {
	userID, _ := GetRequestUserID(c)
	c.Set("IsAdmin", strings.Contains(admins, userID.String()))
}

// AdminUserMiddleware 管理者ユーザーか判定するミドルウェア
func (h *Handlers) AdminUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := GetRequestUserIsAdmin(c)

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

func GetRequestUserIsAdmin(c echo.Context) bool {
	return c.Get("IsAdmin").(bool)
}

func GetRequestUserToken(c echo.Context) (string, error) {
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
