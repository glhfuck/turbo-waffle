package pgrepository

import (
	"fmt"

	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/jmoiron/sqlx"
)

type shortPostgres struct {
	db *sqlx.DB
}

func (sp *shortPostgres) OriginalURL(linkId int) (string, error) {
	tx, err := sp.db.Begin()
	if err != nil {
		return "", err
	}

	selectQuery := `
	SELECT original_URL
	FROM links
	WHERE link_id = $1
	`

	var originalURL string
	row := tx.QueryRow(selectQuery, linkId)
	if err := row.Scan(&originalURL); err != nil {
		tx.Rollback()
		return "", err
	}

	updateQuery := `
	UPDATE links
	SET visits_count = visits_count + 1
	WHERE link_id = $1
	`

	_, err = tx.Exec(updateQuery, linkId)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return originalURL, nil
}

func (sp *shortPostgres) SaveLink(link *domain.Link) (*domain.Link, error) {
	query := `
	INSERT INTO links (owner_id, original_url, creation_date, update_date, visits_count)
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING link_id
	`

	fmt.Println(link.OwnerId)

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

func NewShortPostgres(db *sqlx.DB) *shortPostgres {
	return &shortPostgres{db: db}
}
