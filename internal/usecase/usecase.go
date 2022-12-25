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
	ParseRoute(route string) (string, error)
	ShortURL(originalURL string, userId int) (string, error)
}

type Statistics interface {
	OneStat(userId int, route string) (*domain.Link, error)
	AllStat(userId int) ([]domain.Link, error)
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		Authorization: newAuthUsecase(repo.Authorization),
		Shortener:     newShortUsecase(repo.Shortener),
		Statistics:    newStatUsecase(repo.Statistics),
	}
}
