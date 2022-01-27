package router

import (
	"github.com/gofrs/uuid"
	"github.com/traPtitech/Emoine/repository"
	"google.golang.org/protobuf/proto"
)

type rawMessage struct {
	userID      uuid.UUID
	messageType int
	data        []byte
}

func (s *Streamer) logger(m *rawMessage) error {
	msg := &Message{}
	if err := proto.Unmarshal(m.data, msg); err != nil {
		return err
	}

	payload := msg.GetPayload()
	switch payload.(type) {
	case *Message_State:
		if err := s.stateLogger(msg.GetState()); err != nil {
			return err
		}
	case *Message_Reaction:
		if err := s.reactionLogger(m.userID, msg.GetReaction()); err != nil {
			return err
		}
	case *Message_Comment:
		if err := s.commentLogger(m.userID, msg.GetComment()); err != nil {
			return err
		}
	}
	return nil
}

// TODO: repositoryの型をいい感じにする

func (s *Streamer) stateLogger(data *State) error {
	state := repository.State{Status: data.Status.String(), Info: data.Info}
	if err := s.repo.UpdateState(&state); err != nil {
		return err
	}
	return nil
}

func (s *Streamer) reactionLogger(userID uuid.UUID, data *Reaction) error {
	reaction := repository.Reaction{UserID: userID, PresentationID: int(data.PresentationId), Stamp: int(data.Stamp)}
	if err := s.repo.CreateReaction(&reaction); err != nil {
		return err
	}
	return nil
}

func (s *Streamer) commentLogger(userID uuid.UUID, data *Comment) error {
	comment := repository.CreateComment{UserID: userID, PresentationID: int(data.PresentationId), Text: data.Text}
	if err := s.repo.CreateComment(&comment); err != nil {
		return err
	}
	return nil
}
