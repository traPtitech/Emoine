package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type Token struct {
	Token     string    `db:"token"`
	UserID    uuid.UUID `db:"userId"`
	CreatedAt time.Time `db:"createdAt"`
}

type TokenRepository interface {
	CreateToken(token *Token) error
	IsTokenValid(tokenString string) (bool, error)
	CleanupExpiredTokens() (int64, error)
}
