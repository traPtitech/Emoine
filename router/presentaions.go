package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetPresentations GET /presentations
func (h *Handlers) GetPresentations(c echo.Context) error {
	fmt.Println("GET /presentations")
	return c.NoContent(http.StatusNoContent)
}
