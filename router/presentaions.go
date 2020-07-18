package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetPresentations GET /presentations
func (h *Handlers) GetPresentations(c echo.Context) error {
	presentations, err := h.Repo.GetPresentations()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, presentations)
}
