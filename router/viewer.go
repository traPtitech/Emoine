package router

import (
	"github.com/traPtitech/Emoine/pb"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetViewer GET /viewer
func (h *Handlers) GetViewer(c echo.Context) error {
	res := &pb.Viewer{Count: uint32(h.streamer.ClientsCount())}
	return c.JSON(http.StatusOK, res)
}
