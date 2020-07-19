package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PostState POST /state
func (h *Handlers) PostState(c echo.Context) error {
	isAdmin := getRequestUserIsAdmin(c)
	if !isAdmin {
		return c.NoContent(http.StatusForbidden)
	}

	var newState *State

	state := c.QueryParam("state")
	if state == "next" {
		switch stateData.GetStatus() {
		case Status_pause:
			newState = &State{
				Status: Status_speaking,
				Info:   "",
				PresentationId: stateData.GetPresentationId(),
			}
		case Status_speaking:
			newState = &State{
				Status: Status_reviewing,
				Info:   "レビュー中",
				PresentationId: stateData.GetPresentationId(),
			}
		case Status_reviewing:
			newState = &State{
				Status: Status_pause,
				Info:   "",
				PresentationId: stateData.GetPresentationId(), //次のIDにする
			}
		}
	} else if state == "pause" {
		if stateData.GetStatus() != Status_speaking {
			return c.NoContent(http.StatusBadRequest)
		}
		newState = &State{
			Status: Status_pause,
			Info:   "",
			PresentationId: stateData.GetPresentationId(),
		}
	} else if state == "resume" {
		if stateData.GetStatus() != Status_pause {
			return c.NoContent(http.StatusBadRequest)
		}
		newState = &State{
			Status: Status_speaking,
			Info:   "",
			PresentationId: stateData.GetPresentationId(),
		}
	} else {
		return c.NoContent(http.StatusBadRequest)
	}

	h.stream.SendState(newState)

	return c.NoContent(http.StatusOK)
}
