package usecase

import (
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
)

type Usecase struct {
	Authorization
	Shortener
	Statistics
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		Authorization: newAuthUsecase(repo.Authorization),
		Shortener:     newShortUsecase(repo.Shortener),
		Statistics:    newStatUsecase(repo.Statistics),
	}
}
