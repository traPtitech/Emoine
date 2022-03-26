package repository

// unused
type Reaction struct {
	UserID         string `db:"userId"`
	PresentationID int    `db:"presentationId"`
	Stamp          int    `db:"stamp"`
}

type ReactionStatistics struct {
	PresentationID int     `json:"presentationID"`
	Counts         []Count `json:"counts"`
}

type Count struct {
	Stamp int `db:"stamp" json:"stamp"`
	Count int `db:"COUNT(stamp)" json:"count"`
}

type ReactionRepository interface {
	CreateReaction(userID string, presentationID, stamp int) error
	GetReactionStatistics(id int) (*ReactionStatistics, error)
}
