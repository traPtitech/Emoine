package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/traPtitech/Emoine/pb"
)

var currentState *pb.State

func setDefaultStateData() {
	currentState = &pb.State{
		Status: pb.Status_pause,
		Info:   "準備中",
		// nullと同義
		PresentationId: 0,
	}
}

// PostState POST /state
func (h *Handlers) PostState(c echo.Context) error {
	var newState *pb.State

	state := c.QueryParam("state")
	if state == "next" {
		switch currentState.GetStatus() {
		case pb.Status_pause:
			if currentState.GetPresentationId() == 0 {
				presentation, err := h.repo.GetFirstPresentation()
				if err != nil {
					return err
				}

				newState = &pb.State{
					Status:         pb.Status_speaking,
					Info:           "発表中",
					PresentationId: uint32(presentation.ID),
				}
			} else {
				newState = &pb.State{
					Status:         pb.Status_speaking,
					Info:           "発表中",
					PresentationId: currentState.GetPresentationId(),
				}
			}
		case pb.Status_speaking:
			newState = &pb.State{
				Status:         pb.Status_reviewing,
				Info:           "レビュー中",
				PresentationId: currentState.GetPresentationId(),
			}
		case pb.Status_reviewing:
			presentation, err := h.repo.GetPresentation(int(currentState.GetPresentationId()))
			if err != nil {
				return err
			}

			if presentation.Next.Valid {
				newState = &pb.State{
					Status:         pb.Status_pause,
					Info:           "発表開始前",
					PresentationId: uint32(presentation.Next.Int64), //次のID
				}
			} else {
				newState = &pb.State{
					Status:         pb.Status_pause,
					Info:           "準備中",
					PresentationId: 0, //空のID
				}
			}
		}
	} else if state == "pause" {
		if currentState.GetStatus() != pb.Status_speaking {
			return c.NoContent(http.StatusBadRequest)
		}
		newState = &pb.State{
			Status:         pb.Status_pause,
			Info:           "発表一時中断中",
			PresentationId: currentState.GetPresentationId(),
		}
	} else if state == "resume" {
		if currentState.GetStatus() != pb.Status_pause {
			return c.NoContent(http.StatusBadRequest)
		}
		newState = &pb.State{
			Status:         pb.Status_speaking,
			Info:           "発表中",
			PresentationId: currentState.GetPresentationId(),
		}
	} else {
		return c.NoContent(http.StatusBadRequest)
	}

	h.streamer.SendState(newState)
	currentState = newState

	return c.NoContent(http.StatusOK)
}
