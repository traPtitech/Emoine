package repository

import (
	"time"

	"github.com/traPtitech/Emoine/utils"
)

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
	CreatePresentation(name, speakers, description string) error
	UpdatePresentation(name, speakers, description string, prev, next, presentationID int) error
	GetPresentations() ([]*Presentation, error)
	GetFirstPresentation() (*Presentation, error)
	GetPresentation(presentationID int) (*Presentation, error)
	DeletePresentation(presentationID int) error
}
