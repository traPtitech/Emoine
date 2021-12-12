package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type Token struct {
	Token     string    `db:"token" json:"token"`
	UserID    uuid.UUID `db:"userId" json:"userId"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
}

type TokenRepository interface {
	CreateToken(token *Token) error
	IsTokenValid(tokenString string) (bool, error)
	CleanupExpiredTokens() (int64, error)
}
