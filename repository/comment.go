package repository

import "github.com/google/uuid"

type Comment struct {
	UserID         uuid.UUID `db:"userId" json:"userId"`
	PresentationID int       `db:"presentationId" json:"presentationId"`
	Text           string    `db:"text" json:"text"`
}

type CommentRepository interface {
	CreateComment(userID string, presentationID int, comment string) error
	GetComments(presentationID int) ([]*Comment, error)
}
