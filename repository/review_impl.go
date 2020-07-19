package repository

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func (repo *SqlxRepository) IsExistReview(userID uuid.UUID, presenID int) (bool, error) {
	var count int
	if err := repo.db.Get(&count, "SELECT COUNT(*) FROM `review` WHERE `userId` = ? AND `presentationId` = ? LIMIT 1", userID, presenID); err != nil {
		fmt.Printf("%#v\n", err)
		return false, nil
	}
	return count > 0, nil
}

func (repo *SqlxRepository) CreateReview(review *Review) error {
	_, err := repo.db.Exec("INSERT INTO `review` (`userId`, `presentationId`, `skill`, `artistry`, `idea`, `presentation`) VALUES (?, ?, ?, ?, ?, ?)",
		review.UserId, review.PresentationId, review.Score.Skill, review.Score.Artistry, review.Score.Idea, review.Score.Presentation)
	return err
}

func (repo *SqlxRepository) UpdateReview(review *Review) error {
	fmt.Println(review)
	_, err := repo.db.Exec("UPDATE `review` SET `skill` = ?, `artistry` = ?, `idea` = ?, `presentation` = ? WHERE `userId` = ? AND `presentationId` = ?",
		review.Score.Skill, review.Score.Artistry, review.Score.Idea, review.Score.Presentation, review.UserId, review.PresentationId)
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
