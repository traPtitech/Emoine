package streamer

import (
	"github.com/google/uuid"
	"github.com/traPtitech/Emoine/pb"
	"google.golang.org/protobuf/proto"
)

type rawMessage struct {
	userID      uuid.UUID
	messageType int
	data        []byte
}

func (s *Streamer) logger(m *rawMessage) error {
	msg := &pb.Message{}
	if err := proto.Unmarshal(m.data, msg); err != nil {
		return err
	}

	payload := msg.GetPayload()
	switch payload.(type) {
	case *pb.Message_State:
		if err := s.stateLogger(msg.GetState()); err != nil {
			return err
		}
	case *pb.Message_Reaction:
		if err := s.reactionLogger(m.userID, msg.GetReaction()); err != nil {
			return err
		}
	case *pb.Message_Comment:
		if err := s.commentLogger(m.userID, msg.GetComment()); err != nil {
			return err
		}
	}
	return nil
}

func (s *Streamer) stateLogger(data *pb.State) error {
	if err := s.repo.UpdateState(data.Status.String(), data.Info); err != nil {
		return err
	}
	return nil
}

func (s *Streamer) reactionLogger(userID uuid.UUID, data *pb.Reaction) error {
	if err := s.repo.CreateReaction(userID.String(), int(data.PresentationId), int(data.Stamp)); err != nil {
		return err
	}
	return nil
}

func (s *Streamer) commentLogger(userID uuid.UUID, data *pb.Comment) error {
	if err := s.repo.CreateComment(userID.String(), int(data.PresentationId), data.Text); err != nil {
		return err
	}
	return nil
}
