package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetPresentationReview GET /presentations/:presentationID/review
func (h *Handlers) GetPresentationReview(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err = h.repo.GetPresentation(presentationID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	res, err := h.repo.GetReviewStatistics(presentationID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// PutPresentationReview PUT /presentations/review
func (h *Handlers) PutPresentationReview(c echo.Context) error {
	userID, err := getSession(c)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	posted := make([]int, 3)
	if err := c.Bind(&posted); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	for _, presentationID := range posted {
		_, err := h.repo.GetPresentation(presentationID)
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
	}

	if err := h.repo.DeleteReview(userID.String()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	for _, presentationID := range posted {
		if err := h.repo.CreateReview(userID.String(), presentationID); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusOK)
}
