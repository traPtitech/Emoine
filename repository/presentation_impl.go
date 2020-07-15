package repository

func (repo *SqlxRepository) CreatePresentation(args *CreatePresentationArgs) error {
	_, err := repo.db.Exec("INSERT INTO `presentation` (name, speaker, description) VALUES (?, ?, ?)",
		args.Name, args.Speaker, args.Description)

	return err
}
