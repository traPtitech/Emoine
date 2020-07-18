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
	key      string
	userID   uuid.UUID
	req      *http.Request
	streamer *Streamer
	conn     *websocket.Conn
	open     bool
	send     chan *rawMessage
	sync.RWMutex
}

// 受信待受
func (s *session) listenRead() {
	s.conn.SetReadLimit(maxReadMessageSize)
	_ = s.conn.SetReadDeadline(time.Now().Add(pongWait))
	s.conn.SetPongHandler(func(string) error {
		_ = s.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		t, m, err := s.conn.ReadMessage()
		if err != nil {
			break
		}

		if t == websocket.BinaryMessage {
			_ = s.write(t, m)
		}

		//if t == websocket.BinaryMessage {
		//	// unsupported
		//	_ = s.writeMessage(&rawMessage{t: websocket.CloseMessage, data: websocket.FormatCloseMessage(websocket.CloseUnsupportedData, "binary message is not supported.")})
		//	break
		//}
	}
}

// listenWrite 送信待受
func (s *session) listenWrite() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-s.send:
			if !ok {
				return
			}

			if err := s.write(msg.t, msg.data); err != nil {
				return
			}

			if msg.t == websocket.CloseMessage {
				return
			}

		case <-ticker.C:
			_ = s.write(websocket.PingMessage, []byte{})
		}
	}
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
