package repository

func (repo *SqlxRepository) IsExistReview(userID string, presentationID int) (bool, error) {
	var count int
	if err := repo.db.Get(&count, "SELECT COUNT(*) FROM `review` WHERE `userId` = ? AND `presentationId` = ? LIMIT 1", userID, presentationID); err != nil {
		return false, nil
	}
	return count > 0, nil
}

func (repo *SqlxRepository) CreateReview(userID string, presentationID int) error {
	_, err := repo.db.Exec("INSERT INTO `review` (`userID`, `presentationId`) VALUES (?, ?)", userID, presentationID)
	return err
}

func (repo *SqlxRepository) DeleteReview(userID string) error {
	_, err := repo.db.Exec("DELETE FROM `review` WHERE `userId` = ?", userID)
	return err
}

func (repo *SqlxRepository) GetReviewStatistics(presentationID int) (*ReviewStatistics, error) {
	statistics := ReviewStatistics{}
	rows, err := repo.db.Queryx("SELECT `presentationId`, COUNT(*) FROM `review` WHERE presentationId = ? GROUP BY `presentationId`", presentationID)
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
