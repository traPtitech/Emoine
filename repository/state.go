package repository

type State struct {
	Status string `db:"status" json:"status"`
	Info   string `db:"info" json:"info"`
}

type StateRepository interface {
	UpdateState(*State) error
	GetState() (*State, error)
}
