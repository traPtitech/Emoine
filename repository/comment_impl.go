package repository

import "log"

func (repo *SqlxRepository) CreateComment(userID string, presentationID int, comment string) error {
	_, err := repo.db.Exec("INSERT INTO `comment` (userId, presentationId, text) VALUES (?, ?, ?)", userID, presentationID, comment)
	return err
}

func (repo *SqlxRepository) GetComments(id int) ([]*Comment, error) {
	var comments []*Comment
	rows, err := repo.db.Queryx("SELECT `presentationId`, `text` FROM `comment` WHERE presentationId = ?", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatalf("Failed to close: %v, the original error was %v", cerr, err)
		}
	}()

	for rows.Next() {
		comment := Comment{}
		if err := rows.StructScan(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}
