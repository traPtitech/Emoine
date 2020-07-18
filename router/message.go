package router

import (
	"fmt"

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
		if err := c.stateMsgHandler(m.GetState()); err != nil {
			return nil
		}
	case *Message_Reaction:
		if err := c.reactionMsgHandler(m.GetReaction()); err != nil {
			return err
		}
	case *Message_Comment:
		if err := c.commentMsgHandler(m.GetComment()); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) stateMsgHandler(m *State) error {
	// TODO Validate message
	// アカン
	state := repository.State{string(m.Status), m.Info}
	if err := c.streamer.repo.UpdateState(&state); err != nil {
		return err
	}
	return nil
}

func (c *client) reactionMsgHandler(m *Reaction) error {
	// TODO Validate message
	// カス
	fmt.Println(Stamp_name[int32(m.Stamp)])
	fmt.Println(int(m.Stamp))
	reaction := repository.Reaction{c.userID, int(m.PresentationId), Stamp_name[int32(m.Stamp)]}
	if err := c.streamer.repo.CreateReaction(&reaction); err != nil {
		return err
	}
	return nil
}

func (c *client) commentMsgHandler(m *Comment) error {
	// TODO Validate message
	// ダメ
	comment := repository.Comment{c.userID, int(m.PresentationId), m.Text}
	if err := c.streamer.repo.CreateComment(&comment); err != nil {
		return err
	}
	return nil
}
