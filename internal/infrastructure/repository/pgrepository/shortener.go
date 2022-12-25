package pgrepository

import (
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/jmoiron/sqlx"
)

type shortPostgres struct {
	db *sqlx.DB
}

func NewShortPostgres(db *sqlx.DB) *shortPostgres {
	return &shortPostgres{db: db}
}

func (sp *shortPostgres) GetLink(linkId int) (*domain.Link, error) {
	var link domain.Link

	query := `
	SELECT owner_id, original_URL, creation_date, update_date, visits_count
	FROM links
	WHERE link_id = $1
	`

	err := sp.db.Get(&link, query, linkId)

	if err != nil {
		return nil, err
	}

	link.Id = linkId
	return &link, nil
}

func (sp *shortPostgres) SaveLink(link *domain.Link) (*domain.Link, error) {
	query := `
	INSERT INTO links owner_id, original_url, creation_date, update_date, visits_count
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING link_id
	`

	row := sp.db.QueryRow(
		query,
		link.OwnerId,
		link.OriginalURL,
		link.CreationDate,
		link.UpdateDate,
		link.VisitsCount)

	err := row.Scan(&link.Id)
	if err != nil {
		return nil, err
	}

	return link, nil
}
