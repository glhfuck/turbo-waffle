package repository

import (
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/pgrepository"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Shortener
	Statistics
}

type Authorization interface {
	CreateUser(u domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Shortener interface {
	OriginalURL(linkId int) (string, error)
	SaveLink(link *domain.Link) (*domain.Link, error)
}

type Statistics interface {
	OneLink(userId, linkId int) (*domain.Link, error)
	AllLinks(userId int) ([]domain.Link, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: pgrepository.NewAuthPostgres(db),
		Shortener:     pgrepository.NewShortPostgres(db),
		Statistics:    pgrepository.NewStatPostgres(db),
	}
}
