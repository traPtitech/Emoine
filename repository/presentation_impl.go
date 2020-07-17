package repository

//TODO: read lock
func (repo *SqlxRepository) CreatePresentation(args *CreatePresentationArgs) error {
	var lastId int
	if err := repo.db.Get(&lastId, "SELECT id FROM `presentation` WHERE `next` IS NULL LIMIT 1"); err != nil {
		return err
	}
	if res, err := repo.db.Exec("INSERT INTO `presentation` (name, speaker, description ,prev) VALUES (?, ?, ?, ?)",
		args.Name, args.Speaker, args.Description, lastId); err != nil {
		return err
	} else {
		lastId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		_, err = repo.db.Exec("UPDATE `presentation` SET `next` = ? WHERE `id` = ?", lastId, lastId)
		return err
	}
}

func (repo *SqlxRepository) UpdatePresentation(state *State) error {
	if _, err := repo.db.Exec("UPDATE `presentation` SET `id` = :id, `name` = :name, `speakers` = :speakers, description = :description, prev = :prev, next = :next", state); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `next` = :id WHERE `id` = :prev", state); err != nil {
		return err
	}
	if _, err := repo.db.Exec("UPDATE `presentation` SET `prev` = :id WHERE `id` = :next", state); err != nil {
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
	var presentation *Presentation
	if err := repo.db.Get(&presentation, "SELECT * FROM `presentation` WHERE `id` = ?", id); err != nil {
		return nil, err
	}
	return presentation, nil
}

func (repo *SqlxRepository) DeletePresentation(id int) error {
	_, err := repo.db.Exec("DELETE FROM `presentation` WHERE `id` = ?", id)
	return err
}
