package repository

import "database/sql"

func (repo *SqlxRepository) CreatePresentation(name, speakers, description string) error {
	var lastID sql.NullInt32
	if err := repo.db.Get(&lastID, "SELECT id FROM `presentation` WHERE `next` IS NULL LIMIT 1"); err != nil {
		lastID = sql.NullInt32{Int32: 0, Valid: false}
	}
	if res, err := repo.db.Exec("INSERT INTO `presentation` (name, speakers, description ,prev) VALUES (?, ?, ?, ?)", name, speakers, description, lastID); err != nil {
		return err
	} else {
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		_, err = repo.db.Exec("UPDATE `presentation` SET `next` = ? WHERE `id` = ?", lastInsertID, lastID)
		return err
	}
}

func (repo *SqlxRepository) UpdatePresentation(name, speakers, description string, prev, next, presentationID int) error {
	if _, err := repo.db.Exec("UPDATE `presentation` SET `name` = ?, `speakers` = ?, description = ?, prev = ?, next = ? WHERE `id` = ?",
		name, speakers, description, prev, next, presentationID); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `next` = ? WHERE `id` = ?", presentationID, prev); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `prev` = ? WHERE `id` = ?", presentationID, next); err != nil {
		return err
	}
	return nil
}

func (repo *SqlxRepository) GetPresentations() ([]*Presentation, error) {
	var presentation []*Presentation
	if err := repo.db.Select(&presentation, "SELECT * FROM `presentation`"); err != nil {
		return nil, err
	}
	return presentation, nil
}

func (repo *SqlxRepository) GetFirstPresentation() (*Presentation, error) {
	presentation := Presentation{}
	if err := repo.db.Get(&presentation, "SELECT * FROM `presentation` WHERE `prev` IS NULL LIMIT 1"); err != nil {
		return nil, err
	}
	return &presentation, nil
}

func (repo *SqlxRepository) GetPresentation(presentationID int) (*Presentation, error) {
	presentation := Presentation{}
	if err := repo.db.Get(&presentation, "SELECT * FROM `presentation` WHERE `presentationID` = ? LIMIT 1", presentationID); err != nil {
		return nil, err
	}
	return &presentation, nil
}

func (repo *SqlxRepository) DeletePresentation(presentationID int) error {
	type Order struct {
		Prev sql.NullInt32 `db:"prev"`
		Next sql.NullInt32 `db:"next"`
	}
	order := Order{}
	if err := repo.db.Get(&order, "SELECT `prev`, `next` FROM `presentation` WHERE `presentationID` = ? LIMIT 1", presentationID); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `next` = ? WHERE `presentationID` = ?", order.Next, order.Prev); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `prev` = ? WHERE `presentationID` = ?", order.Prev, order.Next); err != nil {
		return err
	}
	_, err := repo.db.Exec("DELETE FROM `presentation` WHERE `presentationID` = ?", presentationID)
	return err
}
