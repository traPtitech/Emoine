package router

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
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
	req      *http.Request
	conn     *websocket.Conn
	streamer *Streamer
	open     bool
	send     chan *rawMessage
}
func (s *session) writeMessage(msg *rawMessage) error {
	if s.closed() {
		return ErrAlreadyClosed
	}

	select {
	case s.send <- msg:
	default:
		return ErrBufferIsFull
	}
	return nil
}

func (s *session) write(messageType int, data []byte) error {
	_ = s.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return s.conn.WriteMessage(messageType, data)
}

func (s *session) close() {
	if !s.closed() {
		s.Lock()
		s.open = false
		s.conn.Close()
		close(s.send)
		s.Unlock()
	}
}

func (s *session) closed() bool {
	s.RLock()
	defer s.RUnlock()

	return !s.open
}

// Key implements Session interface.
func (s *session) Key() string {
	return s.key
}

// UserID implements Session interface.
func (s *session) UserID() uuid.UUID {
	return s.userID
}