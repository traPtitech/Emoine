package repository

import "log"

func (repo *SqlxRepository) CreateComment(comment *Comment) error {
	_, err := repo.db.Exec("INSERT INTO `comment` (userId, presentationId, text) VALUES (:userId, :presentationId, :text)", comment)
	return err
}

func (repo *SqlxRepository) GetComments(id int) (*[]Comment, error) {
	var comments []Comment
	rows, err := repo.db.Queryx("SELECT `userId`, `presentationId`, `text` FROM `comment` WHERE presentationId = ?", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatalf("Failed to close: %v, the original error was %v", cerr, err)
		}
	}()

	for rows.Next() {
		if err := rows.StructScan(&comments); err != nil {
			return nil, err
		}
	}
	return &comments, nil
}
