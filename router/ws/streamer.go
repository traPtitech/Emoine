package ws

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/FujishigeTemma/Emoine/event"
	"github.com/FujishigeTemma/Emoine/router"
	"github.com/FujishigeTemma/Emoine/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/leandro-lugaresi/hub"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// ErrAlreadyClosed 既に閉じられています
	ErrAlreadyClosed = errors.New("already closed")
	// ErrBufferIsFull 送信バッファが溢れました
	ErrBufferIsFull = errors.New("buffer is full")

	wsConnectionCounter = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "traq",
		Name:      "ws_connections",
	})
)

// Streamer WebSocketストリーマー
type Streamer struct {
	hub        *hub.Hub
	clients    map[*client]struct{}
	register   chan *client
	unregister chan *client
	stop       chan struct{}
	open       bool
	mu         sync.RWMutex
}

// NewStreamer WebSocketストリーマーを生成し起動します
func NewStreamer(hub *hub.Hub) *Streamer {
	h := &Streamer{
		hub:        hub,
		clients:    make(map[*client]struct{}),
		register:   make(chan *client),
		unregister: make(chan *client),
		stop:       make(chan struct{}),
		open:       true,
	}

	go h.run()
	return h
}

func (s *Streamer) run() {
	for {
		select {
		case client := <-s.register:
			s.mu.Lock()
			s.clients[client] = struct{}{}
			s.mu.Unlock()

		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
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
	fmt.Println("aa")
	if s.IsClosed() {
		http.Error(c.Response(), http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	userID, err := router.GetRequestUserID(c)
	if err != nil {
		return
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		return
	}

	client := &client{
		key:      utils.RandAlphabetAndNumberString(20),
		req:      c.Request(),
		conn:     conn,
		open:     true,
		streamer: s,
		send:     make(chan *rawMessage, messageBufferSize),
		userID:   userID,
	}

	s.register <- client
	wsConnectionCounter.Inc()
	s.hub.Publish(hub.Message{
		Name: event.WSConnected,
		Fields: hub.Fields{
			"user_id": client.UserID(),
			"req":     c.Request(),
		},
	})

	go client.writeLoop()
	client.readLoop()

	s.hub.Publish(hub.Message{
		Name: event.WSDisconnected,
		Fields: hub.Fields{
			"user_id": client.UserID(),
			"req":     c.Request(),
		},
	})
	wsConnectionCounter.Dec()
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
