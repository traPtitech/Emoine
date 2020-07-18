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

	res, err := h.Repo.GetReviewStatistics(presentationID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
