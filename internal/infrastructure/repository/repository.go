package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	Shortener
	Statistics
}

type Authorization interface {
}

type Shortener interface {
}

type Statistics interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
