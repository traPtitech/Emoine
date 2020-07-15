package repository

import "time"

type State struct {
	ID        string    `db:"id"`
	Status    string    `db:"status"`
	Info      string    `db:"info"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type StateRepository interface {
	UpdateState(status string, info string) error
	GetState() (*State, error)
}
