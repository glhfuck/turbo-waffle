package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
