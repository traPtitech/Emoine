package repository

import "log"

func (repo *SqlxRepository) CreateComment(comment *Comment) error {
	result, err := repo.db.Exec("INSERT INTO `comment` (userId, presentationId, text) VALUES (?, ?, ?)", comment.UserID, comment.PresentationID, comment.Text)

	if err != nil {
		return err
	}

	res, err := result.LastInsertId()

	if err != nil {
		return err
	}

	comment.ID = int(res)

	return nil
}

func (repo *SqlxRepository) GetComments(id int) ([]*Comment, error) {
	var comments []*Comment
	rows, err := repo.db.Queryx("SELECT `id`, `userId`, `presentationId`, `text` FROM `comment` WHERE presentationId = ?", id)
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
