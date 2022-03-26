package repository

func (repo *SqlxRepository) UpdateState(status, info string) error {
	_, err := repo.db.Exec("UPDATE `state` SET `status` = ?, `info` = ?", status, info)
	return err
}

func (repo *SqlxRepository) GetState() (*State, error) {
	var state State
	if err := repo.db.Get(&state, "SELECT * FROM `state` LIMIT 1"); err != nil {
		return nil, err
	} else {
		return &state, nil
	}
}
