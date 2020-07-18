package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func (repo *SqlxRepository) CreateReaction(reaction *Reaction) error {
	_, err := repo.db.Exec("INSERT INTO `reaction` (userId, presentationId, stamp) VALUES ( ?, ?, ?)", reaction.UserId, reaction.PresentationId, reaction.Stamp)
	return err
}

func (repo *SqlxRepository) GetReactionStatistics(id int) (*ReactionStatistics, error) {
	var statistics ReactionStatistics
	statistics.PresentationId = id

	var rows *sqlx.Rows
	rows, err := repo.db.Queryx("SELECT `stamp`, COUNT(`stamp`) FROM `reaction` WHERE presentationId = ? GROUP BY `stamp`", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatalf("Failed to close: %v, the original error was %v", cerr, err)
		}
	}()

	for rows.Next() {
		if err := rows.StructScan(&statistics.Counts); err != nil {
			return nil, err
		}
	}
	return &statistics, nil
}
