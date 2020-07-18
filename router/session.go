package router

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var sessionCache = cache.New(2*time.Hour, 3*time.Hour)

type userResponse struct {
	ID      uuid.UUID `json:"id"`
	IsAdmin bool      `json:"isAdmin"`
}

// HandleGetUserMe ヘッダー情報からuser情報を取得
// 認証状態を確認
func (h *Handlers) GetUserMe(c echo.Context) error {
	userID, err := getRequestUserID(c)
	if err != nil {
		return unauthorized(err)
	}
	token, err := getRequestUserToken(c)
	if err != nil {
		return unauthorized(err)
	}

	accessToken, ok := sessionCache.Get(userID.String())
	if !ok {
		return unauthorized(errors.New("no session found for user"))
	}
	if accessToken != token {
		return unauthorized(errors.New("invalid access_token"))
	}

	data := &userResponse{
		userID,
		getRequestUserIsAdmin(c),
	}

	return c.JSON(http.StatusOK, data)
}
