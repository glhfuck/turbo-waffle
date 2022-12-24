package usecase

import (
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
)

type Usecase struct {
	Authorization
	Shortener
	Statistics
}

type Authorization interface {
	CreateUser(u domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Shortener interface {
}

type Statistics interface {
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		Authorization: newAuthUsecase(repo.Authorization),
	}
}
