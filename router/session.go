package router

import (
	"errors"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

var sessionCache = cache.New(2*time.Hour, 3*time.Hour)

type userResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	IsAdmin bool      `json:"isAdmin"`
}

// HandleGetUserMe ヘッダー情報からuser情報を取得
// 認証状態を確認
func (h *Handlers) GetUserMe(c echo.Context) error {
	userID, err := GetRequestUserID(c)
	if err != nil {
		return unauthorized(err)
	}
	userName, err := GetRequestUserName(c)
	if err != nil {
		return unauthorized(err)
	}

	token, err := GetRequestUserToken(c)
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
		userName,
		GetRequestUserIsAdmin(c),
	}

	return c.JSON(http.StatusOK, data)
}
