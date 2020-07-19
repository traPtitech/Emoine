package repository

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func (repo *SqlxRepository) IsExistReview(userID uuid.UUID, presenID int) (bool, error) {
	fmt.Println(userID)
	fmt.Println(presenID)
	type ReviewA struct {
		UserId         uuid.UUID `db:"userId"`
		PresentationId int       `db:"presentationId"`
		Skill          int       `db:"skill"`
		Artistry       int       `db:"artistry"`
		Idea           int       `db:"idea"`
		Presentation   int       `db:"presentation"`
	}
	review := ReviewA{}
	if err := repo.db.Get(&review, "SELECT * FROM `review` WHERE `userId` = ? AND `presentationId` = ? LIMIT 1", userID, presenID); err != nil {
		return false, nil
	}
	fmt.Println(review)
	return &review != nil, nil
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
