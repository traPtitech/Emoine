package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

	var newState *State
	switch stateData.GetStatus() {
	case Status_pause:
		// TODO
	case Status_speaking:
		newState = &State{
			Status: Status_reviewing,
			Info:   "レビュー中",
			PresentationId: stateData.GetPresentationId(),
		}
	case Status_reviewing:
		// TODO
	}

	h.stream.SendAll(&Message{
		Payload: &Message_State{
			State: newState,
		},
	})

	return c.NoContent(http.StatusOK)
}
