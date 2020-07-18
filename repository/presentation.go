package repository

import (
	"database/sql"
	"time"
)

type CreatePresentation struct {
	Name        string `db:"name"`
	Speakers    string `db:"speakers"`
	Description string `db:"description"`
}

type Presentation struct {
	ID          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Speakers    sql.NullString `db:"speakers"`
	Description sql.NullString `db:"description"`
	Prev        sql.NullInt32  `db:"prev"`
	Next        sql.NullInt32  `db:"next"`
	CreatedAt   time.Time      `db:"createdAt"`
	UpdatedAt   time.Time      `db:"updatedAt"`
}

type PresentationRepository interface {
	CreatePresentation(presentation *CreatePresentation) error
	UpdatePresentation(presentation *Presentation) error
	GetPresentations() ([]*Presentation, error)
	GetPresentation(id int) (*Presentation, error)
	DeletePresentation(id int) error
}
