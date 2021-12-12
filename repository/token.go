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
	GetTokenUserID(tokenString string) (uuid.UUID, error)
	CleanupExpiredTokens() (int64, error)
}
