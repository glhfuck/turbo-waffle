package pgrepository

import (
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/jmoiron/sqlx"
)

type statPostgres struct {
	db *sqlx.DB
}

func NewStatPostgres(db *sqlx.DB) *statPostgres {
	return &statPostgres{db: db}
}

func (sp *statPostgres) OneLink(userId, linkId int) (*domain.Link, error) {
	var link domain.Link

	query := `
	SELECT link_id, owner_id, original_url, creation_date, update_date, visits_count
	FROM links
	WHERE owner_id=$1 AND link_id=$2
	`
	err := sp.db.Get(&link, query, userId, linkId)

	return &link, err
}

func (sp *statPostgres) AllLinks(userId int) (*domain.Link, error) {
	return nil, nil
}
