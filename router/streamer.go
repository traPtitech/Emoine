package router

import (
	"errors"
	"github.com/leandro-lugaresi/hub"
	"sync"
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
func NewStreamer(hub *hub.Hub) *Streamer {
	h := &Streamer{
		sessions:   make(map[*session]struct{}),
		register:   make(chan *session),
		unregister: make(chan *session),
		stop:       make(chan struct{}),
		open:       true,
	}

	go h.run()
	return h
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
			m := &Message{
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
