package repository

import (
	"github.com/gofrs/uuid"
)

type Comment struct {
	UserId         uuid.UUID `db:"userId"`
	PresentationId int       `db:"presentationId"`
	Text           string    `db:"text"`
}

type CommentRepository interface {
	CreateComment(comment *Comment) error
	GetComments(id int) (*[]Comment, error)
}
