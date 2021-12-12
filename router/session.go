package router

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type userResponse struct {
	ID      uuid.UUID `json:"id"`
	IsAdmin bool      `json:"isAdmin"`
}

// HandleGetUserMe ヘッダー情報からuser情報を取得
// 認証状態を確認
func (h *Handlers) GetUserMe(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return echo.ErrInternalServerError
	}
	if userID == uuid.Nil {
		return echo.ErrUnauthorized
	}
	isAdmin, err := getIsAdmin(c)
	if err != nil {
		return echo.ErrInternalServerError
	}

	data := &userResponse{
		userID,
		isAdmin,
	}

	return c.JSON(http.StatusOK, data)
}
