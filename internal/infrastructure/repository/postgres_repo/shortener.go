package postgres_repo

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

func (au *shortPostgres) SaveLink(link *domain.Link) (*domain.Link, error) {
	query := `
	INSERT INTO links (owner_id, original_URL, creation_date, update_date, visits_count)
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING link_id
	`

	row := au.db.QueryRow(
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
