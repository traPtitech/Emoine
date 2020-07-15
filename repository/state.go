package repository

type State struct {
	Status string `db:"status"`
	Info   string `db:"info"`
}

type StateRepository interface {
	UpdateState(*State) error
	GetState() (*State, error)
}
