package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetViewer GET /viewer
func (h *Handlers) GetViewer(c echo.Context) error {
	res := &Viewer{ Count: uint32(len(h.stream.clients)) }
	return c.JSON(http.StatusOK, res)
}
