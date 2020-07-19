package router

import (
	"errors"
	"net/http"
	"sync"

	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/FujishigeTemma/Emoine/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

var (
	// ErrAlreadyClosed 既に閉じられています
	ErrAlreadyClosed = errors.New("already closed")
	// ErrBufferIsFull 送信バッファが溢れました
	ErrBufferIsFull = errors.New("buffer is full")
)

// やっぱhub必要かも
// Streamer WebSocketストリーマー
type Streamer struct {
	repo       repository.Repository // これキモい
	clients    map[*client]struct{}
	register   chan *client
	unregister chan *client
	stop       chan struct{}
	open       bool
	mu         sync.RWMutex
}

// NewStreamer WebSocketストリーマーを生成し起動します
func NewStreamer(repo repository.Repository) *Streamer {
	s := &Streamer{
		repo:       repo, // これ
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

// SendAll すべてのclientにメッセージを送る
func (s *Streamer) SendAll(m *Message) {
	byteMessage, err := proto.Marshal(m)
	if err != nil {
		return
	}

	for client := range s.clients {
		client.write(websocket.BinaryMessage, byteMessage)
	}
}

var stateData *State

func setDefaultStateData() {
	stateData = &State{
		Status: Status_pause,
		Info:   "準備中",
		// nullと同義
		PresentationId: 0,
	}
}

// SendState すべてのclientに新しいstateを送る
func (s *Streamer) SendState(st *State) {
	s.SendAll(&Message{
		Payload: &Message_State{
			State: st,
		},
	})
	stateData = st
}

// ServeHTTP http.Handlerインターフェイスの実装
func (s *Streamer) ServeHTTP(c echo.Context) {
	if s.IsClosed() {
		http.Error(c.Response(), http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	userID, err := getRequestUserID(c)
	if err != nil {
		return
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		return
	}

	client := &client{
		key:      utils.RandAlphabetAndNumberString(20),
		userID:   userID,
		req:      c.Request(),
		streamer: s,
		conn:     conn,
		open:     true,
		send:     make(chan *rawMessage, messageBufferSize),
	}

	s.register <- client

	m := &Message{
		Payload: &Message_State{
			State: stateData,
		},
	}
	data, err := proto.Marshal(m)
	if err == nil {
		client.write(websocket.BinaryMessage, data)
	}

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
