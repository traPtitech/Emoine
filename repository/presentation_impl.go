package repository

//TODO: read lock
func (repo *SqlxRepository) CreatePresentation(presentation *CreatePresentation) error {
	var lastID int
	if err := repo.db.Get(&lastID, "SELECT id FROM `presentation` WHERE `next` IS NULL LIMIT 1"); err != nil {
		lastID = -1
	}
	if res, err := repo.db.Exec("INSERT INTO `presentation` (name, speakers, description ,prev) VALUES (?, ?, ?, ?)",
		presentation.Name, presentation.Speakers, presentation.Description, lastId); err != nil {
		return err
	} else {
		lastInsertID, err := res.LastInsertID()
		if err != nil {
			return err
		}
		_, err = repo.db.Exec("UPDATE `presentation` SET `next` = ? WHERE `id` = ?", lastInsertId, lastId)
		return err
	}
}

func (repo *SqlxRepository) UpdatePresentation(presentation *Presentation) error {
	if _, err := repo.db.Exec("UPDATE `presentation` SET `name` = ?, `speakers` = ?, description = ?, prev = ?, next = ? WHERE `id` = ?",
		presentation.ID, presentation.Name, presentation.Speakers, presentation.Description, presentation.Prev, presentation.Next); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `next` = ? WHERE `id` = ?", presentation.ID, presentation.Prev); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `prev` = ? WHERE `id` = ?", presentation.ID, presentation.Next); err != nil {
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

func (repo *SqlxRepository) GetPresentation(id int) (*Presentation, error) {
	presentation := Presentation{}
	if err := repo.db.Get(&presentation, "SELECT * FROM `presentation` WHERE `id` = ? LIMIT 1", id); err != nil {
		return nil, err
	}
	return &presentation, nil
}

func (repo *SqlxRepository) DeletePresentation(id int) error {
	_, err := repo.db.Exec("DELETE FROM `presentation` WHERE `id` = ?", id)
	return err
}
