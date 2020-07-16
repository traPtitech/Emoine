package repository

import (
	"github.com/gofrs/uuid"
)

type Review struct {
	UserId         uuid.UUID `db:"userId"`
	PresentationId int       `db:"presentationId"`
	Scores         struct {
		Skill        int `db:"skill"`
		Artistry     int `db:"artistry"`
		Idea         int `db:"idea"`
		Presentation int `db:"presentation"`
	}
}

type ReviewStatistics struct {
	PresentationId  int `db:"presentationId"`
	Count           int `db:"count"`
	AvgSkill        int `db:"skill"`
	AvgArtistry     int `db:"artistry"`
	AvgIdea         int `db:"idea"`
	AvgPresentation int `db:"presentation"`
}

type ReviewRepository interface {
	CreateReview(review *Review) error
	UpdateReview(review *Review) error
	GetReviewStatistics(id int) (*ReviewStatistics, error)
}
