package router

import (
	"github.com/FujishigeTemma/Emoine/repository"
	"github.com/golang/protobuf/proto"
)

type rawMessage struct {
	t    int
	data []byte
}

var stateData []byte

func setDefaultStateData() {
	m := &Message{
		Payload: &Message_State{
			State: &State{
				Status: Status_pause,
				Info:   "準備中",
			},
		},
	}
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	stateData = data
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
		stateData = b
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
	reaction := repository.Reaction{c.userID, int(m.PresentationId), int(m.Stamp)}
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
