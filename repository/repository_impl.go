package repository

import "github.com/jmoiron/sqlx"

// SqlxRepository sqlxリポジトリ実装
type SqlxRepository struct {
	db *sqlx.DB
}

// NewSqlxRepository リポジトリ実装を初期化して生成します
func NewSqlxRepository(db *sqlx.DB) (*SqlxRepository, error) {
	repo := &SqlxRepository{
		db: db,
	}
	return repo, nil
}
