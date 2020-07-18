package router

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	writeWait          = 10 * time.Second
	pongWait           = 60 * time.Second
	pingPeriod         = (pongWait * 9) / 10
	maxReadMessageSize = 1 << 9 // 512B
	messageBufferSize  = 256
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)
