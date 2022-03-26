package repository

type State struct {
	Status string `db:"status" json:"status"`
	Info   string `db:"info" json:"info"`
}

type StateRepository interface {
	UpdateState(status, info string) error
	GetState() (*State, error)
}
