package router

import (
	"github.com/gofrs/uuid"
	"golang.org/x/net/websocket"
	"sync"
)

type Session interface {
	Key() string
	// UserID このセッションのUserID
	UserID() uuid.UUID
}

type session struct {
	key    string
	userID uuid.UUID
	sync.RWMutex
	conn     *websocket.Conn
	streamer *Streamer
	open     bool
	send     chan *Message
}
