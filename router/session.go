package router

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type userResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	IsAdmin bool      `json:"isAdmin"`
}

// HandleGetUserMe ヘッダー情報からuser情報を取得
// 認証状態を確認
func (h *Handlers) GetUserMe(c echo.Context) error {
	userID, err := getRequestUserID(c)
	if err != nil {
		return unauthorized(err)
	}
	userName, err := getRequestUserName(c)
	if err != nil {
		return unauthorized(err)
	}

	data := &userResponse{
		userID,
		userName,
		getRequestUserIsAdmin(c),
	}

	return c.JSON(http.StatusOK, data)
}
