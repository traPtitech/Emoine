package router

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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

// TraQUserMiddleware traQユーザーか判定するミドルウェア
func (h *Handlers) IsTraQUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := getRequestUserID(c)
		if err != nil {
			return echo.ErrInternalServerError
		}
		if userID == uuid.Nil {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// AdminUserMiddleware 管理者ユーザーか判定するミドルウェア
func (h *Handlers) AdminUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin, err := getRequestUserIsAdmin(c)
		if err != nil {
			log.Printf("error: %v", err)
			return echo.ErrInternalServerError
		}
		if !isAdmin {
			log.Printf("error: %v", err)
			return echo.NewHTTPError(http.StatusForbidden, "You are not admin user.")
		}

		return next(c)
	}
}

func getRequestUserID(c echo.Context) (uuid.UUID, error) {
	sess, err := session.Get("e_session", c)
	if err != nil {
		return uuid.Nil, err
	}
	userID, ok := sess.Values["userID"]
	if !ok {
		return uuid.Nil, nil
	}
	return uuid.FromString(userID.(string))
}

func getRequestUserName(c echo.Context) (string, error) {
	sess, err := session.Get("e_session", c)
	if err != nil {
		return "", err
	}
	userName, ok := sess.Values["userName"]
	if !ok {
		return "", nil
	}
	return userName.(string), nil
}

func getRequestUserIsAdmin(c echo.Context) (bool, error) {
	userID, err := getRequestUserID(c)
	if err != nil {
		return false, err
	}
	return strings.Contains(admins, userID.String()), nil
}
