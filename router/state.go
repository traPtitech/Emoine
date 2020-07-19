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
			if stateData.GetPresentationId() == 0 {
				presentation, err := h.Repo.GetFirstPresentation()
				if err != nil {
					return err
				}

				newState = &State{
					Status:         Status_speaking,
					Info:           "",
					PresentationId: uint32(presentation.ID),
				}
			} else {
				newState = &State{
					Status:         Status_speaking,
					Info:           "",
					PresentationId: stateData.GetPresentationId(),
				}
			}
		case Status_speaking:
			newState = &State{
				Status:         Status_reviewing,
				Info:           "レビュー中",
				PresentationId: stateData.GetPresentationId(),
			}
		case Status_reviewing:
			presentation, err := h.Repo.GetPresentation(int(stateData.GetPresentationId()))
			if err != nil {
				return err
			}

			if presentation.Next.Valid {
				newState = &State{
					Status:         Status_pause,
					Info:           "",
					PresentationId: uint32(presentation.Next.Int64), //次のID
				}
			} else {
				newState = &State{
					Status:         Status_pause,
					Info:           "",
					PresentationId: 0, //空のID
				}
			}
		}
	} else if state == "pause" {
		if stateData.GetStatus() != Status_speaking {
			return c.NoContent(http.StatusBadRequest)
		}
		newState = &State{
			Status:         Status_pause,
			Info:           "",
			PresentationId: stateData.GetPresentationId(),
		}
	} else if state == "resume" {
		if stateData.GetStatus() != Status_pause {
			return c.NoContent(http.StatusBadRequest)
		}
		newState = &State{
			Status:         Status_speaking,
			Info:           "",
			PresentationId: stateData.GetPresentationId(),
		}
	} else {
		return c.NoContent(http.StatusBadRequest)
	}

	h.stream.SendState(newState)

	return c.NoContent(http.StatusOK)
}
