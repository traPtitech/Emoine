package router

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type Client interface {
	Key() string
	UserID() uuid.UUID
	ListenRead()
	ListenWrite()
	PushMessage(*rawMessage) error
	IsClosed() bool
	Close() error
}

type client struct {
	key      string
	userID   uuid.UUID
	conn     *websocket.Conn
	receiver *chan *rawMessage
	sender   chan *rawMessage
	wg       *sync.WaitGroup
	active   bool
	sync.RWMutex
}

// ListenRead クライアントからの受信待受
func (c *client) ListenRead(ctx context.Context) {
	c.conn.SetReadLimit(maxReadMessageSize)

	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		log.Printf("error: %v", err)
	}

	c.conn.SetPongHandler(func(string) error {
		err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			log.Printf("error: %v", err)
		}
		return nil
	})

	for {
		t, d, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		*c.receiver <- &rawMessage{c.userID, t, d}
		if ctx.Err() == context.Canceled {
			return
		}
	}
}

// ListenWrite クライアントへの送信待受
func (c *client) ListenWrite(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.Close(); err != nil {
			log.Printf("error: %v", err)
			return
		}
	}()

	for {
		select {
		case m := <-c.sender:
			if err := c.writeMessage(m.messageType, m.data); err != nil {
				break
			}
			if m.messageType == websocket.CloseMessage {
				return
			}
		case <-ticker.C:
			if err := c.writeMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("error: %v", err)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

// PushMessage メッセージを送信キューに追加
func (c *client) PushMessage(m *rawMessage) error {
	if c.IsClosed() {
		return ErrAlreadyClosed
	}
	select {
	case c.sender <- m:
	default:
		return ErrBufferIsFull
	}
	return nil
}

func (c *client) writeMessage(messageType int, data []byte) error {
	err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		log.Printf("error: %v", err)
	}

	return c.conn.WriteMessage(messageType, data)
}

// IsClosed コネクションの接続状態
func (c *client) IsClosed() bool {
	c.RLock()
	defer c.RUnlock()

	return !c.active
}

// IsClosedWithoutLock コネクションの接続状態
func (c *client) IsClosedWithoutLock() bool {
	return !c.active
}

// Close WebSocketコネクションを切断
func (c *client) Close() error {
	c.Lock()
	defer c.Unlock()

	if c.IsClosedWithoutLock() {
		return ErrAlreadyClosed
	}
	if err := c.conn.Close(); err != nil {
		log.Printf("error: %v", err)
	}
	close(c.sender)
	c.active = false

	c.wg.Done()

	return nil
}

// Key クライアントの識別子
func (c *client) Key() string {
	return c.key
}

// UserID ユーザーID
func (c *client) UserID() uuid.UUID {
	return c.userID
}
