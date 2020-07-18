package router

import "github.com/golang/protobuf/proto"

type rawMessage struct {
	t    int
	data []byte
}

func msgHandler(b []byte) error {
	m := &Message{}
	if err := proto.Unmarshal(b, m); err != nil {
		return err
	}

	payload := m.GetPayload()
	switch payload.(type) {
	case *Message_State:

	case *Message_Reaction:

	case *Message_Comment:

	}
	return nil
}
