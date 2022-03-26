package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetMyPresentationReviews GET /presentations/review/me
func (h *Handlers) GetMyPresentationReviews(c echo.Context) error {
	userID, err := getSession(c)
	if err != nil {
		return err
	}

	res, err := h.repo.GetReviews(userID.String())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("failed to get reviews: %s", err))
	}

	return c.JSON(http.StatusOK, res)
}

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
		return err
	}

	posted := make([]int, 3)
	if err := c.Bind(&posted); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if len(posted) > 3 {
		return echo.ErrBadRequest
	}

	for _, presentationID := range posted {
		_, err := h.repo.GetPresentation(presentationID)
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
	}

	if err := h.repo.DeleteReview(userID.String()); err != nil {
		return err
	}

	for _, presentationID := range posted {
		if err := h.repo.CreateReview(userID.String(), presentationID); err != nil {
			return err
		}
	}

	return c.NoContent(http.StatusOK)
}
