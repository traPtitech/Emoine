package repository

import "time"

type Presentation struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Speaker     string    `db:"speaker"`
	Description string    `db:"description"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
