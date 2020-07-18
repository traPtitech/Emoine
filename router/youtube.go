package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Live struct {
	ID string `json:"liveId"`
}

var live = &Live{"MXMCe6J3YA8"}

func (h *Handlers) GetLiveID(c echo.Context) error {
	return c.JSON(http.StatusOK, live)
}

func (h *Handlers) PutLiveID(c echo.Context) error {
	newLive := Live{}
	if err := c.Bind(&newLive); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	live.ID = newLive.ID
	return c.NoContent(http.StatusOK)
}
