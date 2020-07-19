package ws

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type Client interface {
	Key() string
	// UserID このセッションのUserID
	UserID() uuid.UUID
}

type client struct {
	key      string
	userID   uuid.UUID
	req      *http.Request
	streamer *Streamer
	conn     *websocket.Conn
	open     bool
	send     chan *rawMessage
	sync.RWMutex
}



func (s *client) readLoop() {
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

		if t == websocket.TextMessage {
			s.commandHandler(string(m))
		}

		if t == websocket.BinaryMessage {
			// unsupported
			_ = s.writeMessage(&rawMessage{t: websocket.CloseMessage, data: websocket.FormatCloseMessage(websocket.CloseUnsupportedData, "binary message is not supported.")})
			break
		}
	}
}

func (s *client) writeLoop() {
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

// listenRead クライアントからの受信待受
func (c *client) listenRead() {
	c.conn.SetReadLimit(maxReadMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		t, m, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		// カス
		// if err := c.MsgHandler(m); err != nil {
		// 	break
		// }

		if t == websocket.BinaryMessage {
			for client := range c.streamer.clients {
				_ = client.write(t, m)
			}
		}
	}
}

// listenWrite クライアントへの送信待受
func (c *client) listenWrite() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				return
			}

			if err := c.write(msg.t, msg.data); err != nil {
				return
			}

			if msg.t == websocket.CloseMessage {
				return
			}

		case <-ticker.C:
			_ = c.write(websocket.PingMessage, []byte{})
		}
	}
}

func (c *client) writeMessage(msg *rawMessage) error {
	if c.closed() {
		return ErrAlreadyClosed
	}

	select {
	case c.send <- msg:
	default:
		return ErrBufferIsFull
	}
	return nil
}

func (c *client) write(messageType int, data []byte) error {
	_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(messageType, data)
}

func (c *client) close() {
	if !c.closed() {
		c.Lock()
		c.open = false
		c.conn.Close()
		close(c.send)
		c.Unlock()
	}
}

func (c *client) closed() bool {
	c.RLock()
	defer c.RUnlock()

	return !c.open
}

// Key implements Client interface.
func (c *client) Key() string {
	return c.key
}

// UserID implements Client interface.
func (c *client) UserID() uuid.UUID {
	return c.userID
}
