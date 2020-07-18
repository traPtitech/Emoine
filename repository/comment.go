package repository

import (
	"github.com/gofrs/uuid"
)

type Comment struct {
	UserID         uuid.UUID `db:"userId"`
	PresentationID int       `db:"presentationId"`
	Text           string    `db:"text"`
}

type CommentRepository interface {
	CreateComment(comment *Comment) error
	GetComments(id int) (*[]Comment, error)
}
