package repository

import (
	"time"

	"github.com/FujishigeTemma/Emoine/utils"
)

type CreatePresentation struct {
	Name        string `db:"name"`
	Speakers    string `db:"speakers"`
	Description string `db:"description"`
}

type Presentation struct {
	ID          int          `db:"id" json:"id"`
	Name        utils.String `db:"name" json:"name"`
	Speakers    utils.String `db:"speakers" json:"speakers"`
	Description utils.String `db:"description" json:"description"`
	Prev        utils.Int    `db:"prev" json:"prev"`
	Next        utils.Int    `db:"next" json:"next"`
	CreatedAt   time.Time    `db:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time    `db:"updatedAt" json:"updatedAt"`
}

type PresentationRepository interface {
	CreatePresentation(presentation *CreatePresentation) error
	UpdatePresentation(presentation *Presentation) error
	GetPresentations() ([]*Presentation, error)
	GetFirstPresentation() (*Presentation, error)
	GetPresentation(id int) (*Presentation, error)
	DeletePresentation(id int) error
}
