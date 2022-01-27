package repository

import (
	"github.com/gofrs/uuid"
)

type Comment struct {
	ID             int       `db:"id" json:"id"`
	UserID         uuid.UUID `db:"userId" json:"userId"`
	PresentationID int       `db:"presentationId" json:"presentationId"`
	Text           string    `db:"text" json:"text"`
}

type CommentRepository interface {
	CreateComment(comment *Comment) error
	GetComments(id int) ([]*Comment, error)
}
