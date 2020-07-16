package repository

func (repo *SqlxRepository) CreateReview(review *Review) error {
	_, err := repo.db.Exec("INSERT INTO `review` (userId, presentationId, skill, artistry, idea, presentation) VALUES (:userId, :presentationId, :skill, :artistry, :idea, :presentation)", review)
	return err
}

func (repo *SqlxRepository) UpdateReview(review *Review) error {
	_, err := repo.db.Exec("UPDATE `review` SET userId = :userId, presentationId = :presentationId, skill = :skill, artistry = :artistry, idea = :idea, presentation = :presentation", review)
	return err
}

func (repo *SqlxRepository) GetReviewStatistics(id int) (*ReviewStatistics, error) {
	var statistics ReviewStatistics
	row := repo.db.QueryRow("SELECT `presentationId`, COUNT(*), AVG(`skill`), AVG(`artistry`), AVG(`idea`), AVG(`presentation`) FROM `reaction` WHERE presentationId = ? GROUP BY `presentationId`", id)
	if err := row.Scan(&statistics); err != nil {
		return nil, err
	}
	return &statistics, nil
}
