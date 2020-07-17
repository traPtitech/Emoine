package repository

type CreatePresentation struct {
	Name        string    `db:"name"`
	Speaker     string    `db:"speaker"`
	Description string    `db:"description"`
}

type Presentation struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Speaker     string    `db:"speaker"`
	Description string    `db:"description"`
	Prev        int       `db:"prev"`
	Next        int       `db:"next"`
}

type PresentationRepository interface {
	CreatePresentation(presentation *CreatePresentation) error
	UpdatePresentation(presentation *Presentation) error
	GetPresentations() ([]*Presentation, error)
	GetPresentation(id int) (*Presentation, error)
	DeletePresentation(id int) error
}