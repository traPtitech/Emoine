package router

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Emoine/repository"
	"github.com/traPtitech/Emoine/utils"
	"google.golang.org/protobuf/proto"
)

var (
	// ErrAlreadyClosed 既に閉じられています
	ErrAlreadyClosed = errors.New("already IsClosed")
	// ErrBufferIsFull 送信バッファが溢れました
	ErrBufferIsFull = errors.New("buffer is full")
)

// Streamer WebSocketストリーマー
type Streamer struct {
	repo          repository.Repository
	clients       map[string]*client
	registry      chan *client
	messageBuffer chan *rawMessage
	active        bool
	sync.RWMutex
}

// NewStreamer WebSocketストリーマーを生成し起動します
func NewStreamer(repo repository.Repository) *Streamer {
	s := &Streamer{
		repo:          repo,
		clients:       make(map[string]*client),
		registry:      make(chan *client),
		messageBuffer: make(chan *rawMessage),
		active:        true,
	}

	go s.run()
	return s
}

func (s *Streamer) run() {
	for {
		select {
		case client := <-s.registry:
			if client.active {
				s.clients[client.Key()] = client
			} else {
				if _, ok := s.clients[client.Key()]; !ok {
					delete(s.clients, client.Key())
				}
			}
		case m := <-s.messageBuffer:
			s.logger(m)
			s.SendAll(m)
		}
	}
}

// SendAll すべてのclientにメッセージを送る
func (s *Streamer) SendAll(m *rawMessage) {
	for _, client := range s.clients {
		if err := client.PushMessage(m); err != nil {
			log.Printf("error: %v", err)
		}
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
	msg := &Message{
		Payload: &Message_State{
			State: st,
		},
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	for _, client := range s.clients {
		m := &rawMessage{client.UserID(), websocket.BinaryMessage, data}
		if err := client.PushMessage(m); err != nil {
			log.Printf("error: %v", err)
		}
	}
	stateData = st
}

// ServeHTTP GET /ws
func (s *Streamer) ServeHTTP(c echo.Context) error {
	if s.IsClosed() {
		return echo.ErrServiceUnavailable
	}
	userID, err := getRequestUserID(c)
	if err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())
	if err != nil {
		log.Printf("error: %v", err)
		return echo.ErrInternalServerError
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &client{
		key:      utils.RandAlphabetAndNumberString(20),
		userID:   userID,
		conn:     conn,
		receiver: &s.messageBuffer,
		sender:   make(chan *rawMessage, messageBufferSize),
		wg:       &wg,
		active:   true,
	}

	s.registry <- client
	defer func() {
		if !client.IsClosed() {
			if err := client.Close(); err != nil {
				log.Printf("error: %v", err)
			}
		}
		s.registry <- client
	}()

	wg.Add(1)
	go client.ListenWrite(ctx)
	go client.ListenRead(ctx)

	msg := &Message{
		Payload: &Message_State{
			State: stateData,
		},
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("error: %v", err)
	}
	m := &rawMessage{client.UserID(), websocket.BinaryMessage, data}

	if err := client.PushMessage(m); err != nil {
		log.Printf("error: %v", err)
	}

	msg = &Message{
		Payload: &Message_Viewer{
			&Viewer{ Count: uint32(len(s.clients)) },
		},
	}
	data, err = proto.Marshal(msg)
	if err != nil {
		log.Printf("error: %v", err)
	}
	m = &rawMessage{client.UserID(), websocket.BinaryMessage, data}

	if err := client.PushMessage(m); err != nil {
		log.Printf("error: %v", err)
	}

	wg.Wait()

	return c.NoContent(http.StatusOK)
}

// IsClosed ストリーマーが停止しているかどうか
func (s *Streamer) IsClosed() bool {
	s.RLock()
	defer s.RUnlock()

	return !s.active
}

// Close ストリーマーを停止します
func (s *Streamer) Close() error {
	s.Lock()
	defer s.Unlock()

	if s.IsClosed() {
		return ErrAlreadyClosed
	}

	m := &rawMessage{
		messageType: websocket.CloseMessage,
		data:        websocket.FormatCloseMessage(websocket.CloseServiceRestart, "Server is stopping..."),
	}
	for _, client := range s.clients {
		if err := client.PushMessage(m); err != nil {
			log.Printf("error: %v", err)
		}
		delete(s.clients, client.Key())
		if err := client.Close(); err != nil {
			log.Printf("error: %v", err)
		}
	}
	s.active = false

	return nil
}
