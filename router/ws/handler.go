package ws

import (
	// "fmt"
	// "github.com/gofrs/uuid"
	// "github.com/gorilla/websocket"
	"strings"
)

func (s *client) commandHandler(cmd string) {
	args := strings.Split(strings.TrimSpace(cmd), ":")

	switch strings.ToLower(args[0]) {
	case "timeline_streaming":

	default:
		// 不明なコマンド
		// s.sendErrorMessage(fmt.Sprintf("unknown command: %s", cmd))
	}
}

// func (s *client) sendErrorMessage(error string) {
// 	_ = s.writeMessage(&rawMessage{
// 		t:    websocket.TextMessage,
// 		data: makeMessage("ERROR", error).toJSON(),
// 	})
// }
