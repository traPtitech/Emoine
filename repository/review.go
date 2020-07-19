package repository

import (
	"github.com/gofrs/uuid"
)

type Review struct {
	UserId         uuid.UUID `db:"userId"`
	PresentationId int       `db:"presentationId"`
	Score
}

type Score struct {
	Skill        int `db:"skill"`
	Artistry     int `db:"artistry"`
	Idea         int `db:"idea"`
	Presentation int `db:"presentation"`
}

type ReviewStatistics struct {
	PresentationId  int `db:"presentationId" json:"presentationId"`
	Count           int `db:"COUNT(*)" json:"count"`
	AvgSkill        float64 `db:"AVG(skill)" json:"avgSkill"`
	AvgArtistry     float64 `db:"AVG(artistry)" json:"avgArtistry"`
	AvgIdea         float64 `db:"AVG(idea)" json:"avgIdea"`
	AvgPresentation float64 `db:"AVG(presentation)" json:"avgPresentation"`
}

type ReviewRepository interface {
	CreateReview(review *Review) error
	UpdateReview(review *Review) error
	GetReviewStatistics(id int) (*ReviewStatistics, error)
}
