package router

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"sync"
	"unsafe"
)

var (
	// ErrAlreadyClosed 既に閉じられています
	ErrAlreadyClosed = errors.New("already closed")
	// ErrBufferIsFull 送信バッファが溢れました
	ErrBufferIsFull = errors.New("buffer is full")
)

// Streamer WebSocketストリーマー
type Streamer struct {
	clients    map[*client]struct{}
	register   chan *client
	unregister chan *client
	stop       chan struct{}
	open       bool
	mu         sync.RWMutex
}

// NewStreamer WebSocketストリーマーを生成し起動します
func NewStreamer() *Streamer {
	s := &Streamer{
		clients:    make(map[*client]struct{}),
		register:   make(chan *client),
		unregister: make(chan *client),
		stop:       make(chan struct{}),
		open:       true,
	}

	go s.run()
	return s
}

func (s *Streamer) run() {
	for {
		select {
		case client := <-s.register:
			s.mu.Lock()
			s.clients[client] = struct{}{}
			s.mu.Unlock()

		case client := <-s.unregister:
			if _, ok := s.clients[client]; !ok {
				s.mu.Lock()
				delete(s.clients, client)
				s.mu.Unlock()
			}

		case <-s.stop:
			s.mu.Lock()
			m := &rawMessage{
				t:    websocket.CloseMessage,
				data: websocket.FormatCloseMessage(websocket.CloseServiceRestart, "Server is stopping..."),
			}
			for client := range s.clients {
				_ = client.writeMessage(m)
				delete(s.clients, client)
				client.close()
			}
			s.open = false
			s.mu.Unlock()
			return
		}
	}
}

// ServeHTTP http.Handlerインターフェイスの実装
func (s *Streamer) ServeHTTP(c echo.Context) {
	if s.IsClosed() {
		http.Error(c.Response(), http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		return
	}

	client := &client{
		key:      randomAlphaNumeric(20),
		userID:   c.Request().Context().Value("userId").(uuid.UUID),
		req:      c.Request(),
		streamer: s,
		conn:     conn,
		open:     true,
		send:     make(chan *rawMessage, messageBufferSize),
	}

	s.register <- client

	go client.listenWrite()
	client.listenRead()

	s.unregister <- client
	client.close()
}

// IsClosed ストリーマーが停止しているかどうか
func (s *Streamer) IsClosed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return !s.open
}

// Close ストリーマーを停止します
func (s *Streamer) Close() error {
	if s.IsClosed() {
		return ErrAlreadyClosed
	}
	s.stop <- struct{}{}
	return nil
}

const (
	rs6Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1<<rs6LetterIdxBits - 1
	rs6LetterIdxMax  = 63 / rs6LetterIdxBits
)

// randomAlphaNumeric 指定した文字数のランダム英数字文字列を生成します
// この関数はmath/randが生成する擬似乱数を使用します
func randomAlphaNumeric(n int) string {
	b := make([]byte, n)
	cache, remain := rand.Int63(), rs6LetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
