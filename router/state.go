package router

import (
	"net/http"

	"github.com/FujishigeTemma/Emoine/event"
	"github.com/labstack/echo/v4"
	"github.com/leandro-lugaresi/hub"
)

type PostStateStruct struct {
	State string `json:"state" db:"status"`
}

// PostState POST /state
func (h *Handlers) PostState(c echo.Context) error {
	posted := PostStateStruct{}
	if err := c.Bind(&posted); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	isAdmin := getRequestUserIsAdmin(c)
	if !isAdmin {
		return c.NoContent(http.StatusForbidden)
	}
	h.stream.Publish(hub.Message{
		Name: event.StateUpdated,
		Fields: hub.Fields{
			"state": posted.State,
		},
	})
	return c.NoContent(http.StatusOK)
}
