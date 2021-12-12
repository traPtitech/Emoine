package repository

import "database/sql"

func (repo *SqlxRepository) CreateToken(token *Token) error {
	_, err := repo.db.Exec("INSERT INTO `token` (token, userId) VALUES (?, ?)", token.Token, token.UserID)
	return err
}

func (repo *SqlxRepository) IsTokenValid(tokenString string) (bool, error) {
	err := repo.db.Select("SELECT 1 FROM `token` WHERE `token` = ? AND `created_at` >= NOW() - INTERVAL 1 DAY", tokenString)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *SqlxRepository) CleanupExpiredTokens() (int64, error) {
	res, err := repo.db.Exec("DELETE FROM `token` WHERE `created_at` < NOW() - INTERVAL 1 DAY")
	if err != nil {
		return 0, err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}
