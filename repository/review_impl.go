package repository

func (repo *SqlxRepository) CreateReview(review *Review) error {
	_, err := repo.db.Exec("INSERT INTO `review` (`userId`, `presentationId`, `skill`, `artistry`, `idea`, `presentation`) VALUES (?, ?, ?, ?, ?, ?)",
		review.UserId, review.PresentationId, review.Score.Skill, review.Score.Artistry, review.Score.Idea, review.Score.Presentation)
	return err
}

func (repo *SqlxRepository) UpdateReview(review *Review) error {
	_, err := repo.db.Exec("UPDATE `review` SET `presentationId` = ?, `skill` = ?, `artistry` = ?, `idea` = ?, `presentation` = ? WHERE `userId` = ?",
		review.PresentationId, review.Score.Skill, review.Score.Artistry, review.Score.Idea, review.Score.Presentation, review.UserId)
	return err
}

func (repo *SqlxRepository) GetReviewStatistics(id int) (*ReviewStatistics, error) {
	statistics := ReviewStatistics{}
	// QueryRowsでexpected 6 destination arguments in Scan, not 1でバグる
	rows, err := repo.db.Queryx("SELECT `presentationId`, COUNT(*), AVG(skill), AVG(artistry), AVG(idea), AVG(presentation) FROM `review` WHERE presentationId = ? GROUP BY `presentationId`", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.StructScan(&statistics); err != nil {
			return nil, err
		}
	}
	return &statistics, nil
}
