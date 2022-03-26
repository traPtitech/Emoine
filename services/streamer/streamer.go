package streamer

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/traPtitech/Emoine/repository"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	"github.com/traPtitech/Emoine/pb"
	"github.com/traPtitech/Emoine/utils"
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
	rwm           sync.RWMutex
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
				delete(s.clients, client.Key())
			}

			m, err := getViewerMessage(client.UserID(), len(s.clients))
			if err != nil {
				log.Printf("error: %v", err)
				break
			}
			s.SendAll(m)
		case m := <-s.messageBuffer:
			err := s.logger(m)
			if err != nil {
				log.Printf("error: %v", err)
			}

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

func getViewerMessage(userID uuid.UUID, length int) (*rawMessage, error) {
	msg := &pb.Message{
		Payload: &pb.Message_Viewer{
			Viewer: &pb.Viewer{Count: uint32(length)},
		},
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	m := &rawMessage{userID, websocket.BinaryMessage, data}
	return m, nil
}

// SendState すべてのclientに新しいstateを送る
func (s *Streamer) SendState(st *pb.State) {
	msg := &pb.Message{
		Payload: &pb.Message_State{
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
}

// NewClient 新規クライアントを初期化・登録します
func (s *Streamer) NewClient(conn *websocket.Conn, currentState *pb.State) error {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &client{
		key:      utils.RandAlphabetAndNumberString(20),
		userID:   uuid.New(),
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

	msg := &pb.Message{
		Payload: &pb.Message_State{
			State: currentState,
		},
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	m := &rawMessage{client.UserID(), websocket.BinaryMessage, data}

	if err := client.PushMessage(m); err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func (s *Streamer) ClientsCount() int {
	return len(s.clients)
}

// IsClosed ストリーマーが停止しているかどうか
func (s *Streamer) IsClosed() bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	return !s.active
}

// IsClosedWithoutLock ストリーマーが停止しているかどうか
func (s *Streamer) IsClosedWithoutLock() bool {
	return !s.active
}

// Close ストリーマーを停止します
func (s *Streamer) Close() error {
	if s.IsClosedWithoutLock() {
		return ErrAlreadyClosed
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()

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
