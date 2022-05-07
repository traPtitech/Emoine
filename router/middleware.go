package router

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Emoine/utils"
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

const userIDContextKey = "Emoine_UserID"
const adminContextKey = "Emoine_Admin"

func (h *Handlers) TraQUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, isByToken, err := h.getRequestUserID(c)
		if err != nil {
			return err
		}

		isAdmin := strings.Contains(admins, userID.String())

		c.Set(userIDContextKey, userID)
		c.Set(adminContextKey, isAdmin && !isByToken) // トークンでの認証の場合はadminにしない

		return next(c)
	}
}

// TraQUserMiddleware traQユーザーか判定するミドルウェア
func (h *Handlers) IsTraQUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := getUserID(c)
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
func (h *Handlers) IsAdminUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin, err := getIsAdmin(c)
		if err != nil {
			return echo.ErrInternalServerError
		}
		if !isAdmin {
			return echo.NewHTTPError(http.StatusForbidden, "You are not admin user.")
		}
		return next(c)
	}
}

func (h *Handlers) getRequestUserIDFromCookie(c echo.Context) (uuid.UUID, error) {
	sess, err := session.Get("emoine_session_v2", c)
	if err != nil {
		return uuid.Nil, err
	}
	userID, ok := sess.Values["userID"]
	if !ok {
		return uuid.Nil, nil
	}
	return uuid.FromString(userID.(string))
}

func (h *Handlers) getRequestUserIDFromToken(c echo.Context) (uuid.UUID, error) {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" {
		return uuid.Nil, nil
	}
	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return uuid.Nil, errors.New("Invalid authorization type.")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	if token == "" {
		return uuid.Nil, nil
	}

	if err := utils.CheckTokenPrefix(token); err != nil {
		return uuid.Nil, err
	}

	userID, err := h.Repo.GetTokenUserID(token)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

// 二つ目の返り値のboolはtokenでの認証かどうか
func (h *Handlers) getRequestUserID(c echo.Context) (uuid.UUID, bool, error) {
	var errs *multierror.Error

	userID, err1 := h.getRequestUserIDFromCookie(c)
	if err1 == nil && userID != uuid.Nil {
		return userID, false, nil
	}
	errs = multierror.Append(errs, err1)

	userID, err2 := h.getRequestUserIDFromToken(c)
	if err2 == nil && userID != uuid.Nil {
		return userID, true, nil
	}
	errs = multierror.Append(errs, err2)

	if len(errs.Errors) == 0 {
		return uuid.Nil, false, nil
	}
	return uuid.Nil, false, echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("Failed to get userID: %w", errs))
}

func getUserID(c echo.Context) (uuid.UUID, error) {
	userIDI := c.Get(userIDContextKey)
	if userIDI == nil {
		return uuid.Nil, echo.ErrInternalServerError
	}
	return userIDI.(uuid.UUID), nil
}

func getIsAdmin(c echo.Context) (bool, error) {
	isAdminI := c.Get(adminContextKey)
	if isAdminI == nil {
		return false, echo.ErrInternalServerError
	}
	return isAdminI.(bool), nil
}
