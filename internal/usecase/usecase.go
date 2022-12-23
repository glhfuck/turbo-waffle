package usecase

import "github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"

type Usecase struct {
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

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{}
}
