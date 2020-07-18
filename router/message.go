package router

import (
	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/golang/protobuf/proto"
)

type rawMessage struct {
	t    int
	data []byte
}

func (c *client) MsgHandler(b []byte) error {
	m := &Message{}
	if err := proto.Unmarshal(b, m); err != nil {
		return err
	}

	payload := m.GetPayload()
	switch payload.(type) {
	case *Message_State:

	case *Message_Reaction:
		if err := c.reactionMsgHandler(m.GetReaction()); err != nil {
			return err
		}
	case *Message_Comment:

	}
	return nil
}

func (c *client) reactionMsgHandler(m *Reaction) error {
	// TODO Validate message
	// カス
	reaction := repository.Reaction{c.userID, int(m.PresentationId), string(m.Stamp)}
	if err := c.streamer.repo.CreateReaction(&reaction); err != nil {
		return err
	}
	return nil
}
