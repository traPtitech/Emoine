package router

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

// ConnectWebSocket GET /ws
func (h *Handlers) ConnectWebSocket(c echo.Context) error {
	if h.streamer.IsClosed() {
		return echo.ErrServiceUnavailable
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}

	userID, err := getSession(c)
	if err != nil {
		log.Printf("error: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := h.streamer.NewClient(conn, currentState, userID); err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}

	return nil
}
