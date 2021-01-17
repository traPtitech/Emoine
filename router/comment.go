package router

import (
	"net/http"
	"strconv"

	"github.com/traPtitech/Emoine/repository"
	"github.com/labstack/echo/v4"
)

// GetPresentationComments GET /presentations/:presentationID/comments
func (h *Handlers) GetPresentationComments(c echo.Context) error {
	presentationID, err := strconv.Atoi(c.Param("presentationID"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	res, err := h.Repo.GetComments(presentationID)
	if err != nil {
		return err
	}
	if res == nil {
		return c.JSON(http.StatusOK, []repository.Comment{})
	}
	return c.JSON(http.StatusOK, res)
}
