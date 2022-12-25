package usecase

import (
	"errors"
	"strconv"

	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
)

type statUsecase struct {
	repo repository.Statistics
}

func newStatUsecase(repo repository.Statistics) *statUsecase {
	return &statUsecase{repo: repo}
}

func (su *statUsecase) OneStat(userId int, route string) (*domain.Link, error) {
	linkId, err := strconv.ParseInt(route, 36, 32)

	if err != nil {
		return nil, errors.New("can't parse route")
	}

	link, err := su.repo.OneLink(userId, int(linkId))

	return link, err
}

func (su *statUsecase) AllStat(userId int) ([]domain.Link, error) {
	return nil, nil
}
