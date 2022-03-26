package repository

type ReviewStatistics struct {
	PresentationID int `db:"presentationId" json:"presentationId"`
	Count          int `db:"COUNT(*)" json:"count"`
}

type ReviewRepository interface {
	IsExistReview(userID string, presentationID int) (bool, error)
	CreateReview(userID string, presentationID int) error
	DeleteReview(userID string) error
	GetReviews(userID string) ([]int, error)
	GetReviewStatistics(presentationID int) (*ReviewStatistics, error)
}
