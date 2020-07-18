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
	sessions   map[*session]struct{}
	register   chan *session
	unregister chan *session
	stop       chan struct{}
	open       bool
	mu         sync.RWMutex
}

// NewStreamer WebSocketストリーマーを生成し起動します
func NewStreamer() *Streamer {
	s := &Streamer{
		sessions:   make(map[*session]struct{}),
		register:   make(chan *session),
		unregister: make(chan *session),
		stop:       make(chan struct{}),
		open:       true,
	}

	go s.run()
	return s
}

func (s *Streamer) run() {
	for {
		select {
		case session := <-s.register:
			s.mu.Lock()
			s.sessions[session] = struct{}{}
			s.mu.Unlock()

		case session := <-s.unregister:
			if _, ok := s.sessions[session]; ok {
				s.mu.Lock()
				delete(s.sessions, session)
				s.mu.Unlock()
			}

		case <-s.stop:
			s.mu.Lock()
			m := &rawMessage{
				t:    websocket.CloseMessage,
				data: websocket.FormatCloseMessage(websocket.CloseServiceRestart, "Server is stopping..."),
			}
			for session := range s.sessions {
				_ = session.writeMessage(m)
				delete(s.sessions, session)
				session.close()
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

	session := &session{
		key:      randomAlphaNumeric(20),
		userID:   c.Request().Context().Value("userId").(uuid.UUID),
		req:      c.Request(),
		streamer: s,
		conn:     conn,
		open:     true,
		send:     make(chan *rawMessage, messageBufferSize),
	}

	s.register <- session

	go session.listenWrite()
	session.listenRead()

	s.unregister <- session
	session.close()
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
