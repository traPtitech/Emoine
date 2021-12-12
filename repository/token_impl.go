package repository

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func (repo *SqlxRepository) CreateToken(token *Token) error {
	_, err := repo.db.Exec("INSERT INTO `token` (token, userId) VALUES (?, ?)", token.Token, token.UserID)
	return err
}

func (repo *SqlxRepository) GetTokenUserID(tokenString string) (uuid.UUID, error) {
	var userID uuid.UUID

	err := repo.db.Get(&userID, "SELECT `userID` FROM `token` WHERE `token` = ? AND `createdAt` >= NOW() - INTERVAL 1 DAY", tokenString)
	if err == sql.ErrNoRows {
		return uuid.Nil, nil
	} else if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

func (repo *SqlxRepository) CleanupExpiredTokens() (int64, error) {
	res, err := repo.db.Exec("DELETE FROM `token` WHERE `createdAt` < NOW() - INTERVAL 1 DAY")
	if err != nil {
		return 0, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}
