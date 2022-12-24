package repository

import (
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/postgres_repo"
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
}

type Statistics interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres_repo.NewAuthPostgres(db),
	}
}
