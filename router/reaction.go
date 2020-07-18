package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetPresentationReaction GET /presentations/:presentationID/reaction
func (h *Handlers) GetPresentationReaction(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	res, err := h.Repo.GetReactionStatistics(presentationID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
