package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// GetPresentations GET /presentaions
func (h *Handlers) GetPresentations(c echo.Context) error {
	fmt.Println("GET /presentaions")
	return c.NoContent(http.StatusNoContent)
}
