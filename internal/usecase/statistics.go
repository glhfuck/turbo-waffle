package usecase

import (
	"errors"
	"strconv"

	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
)

type Statistics interface {
	OneStat(userId int, route string) (*Stat, error)
	AllStat(userId int) ([]Stat, error)
}

type statUsecase struct {
	repo repository.Statistics
}

type Stat struct {
	Route string `json:"route"`
	Link  *domain.Link
}

func (su *statUsecase) OneStat(userId int, route string) (*Stat, error) {
	linkId, err := strconv.ParseInt(route, 36, 32)

	if err != nil {
		return nil, errors.New("can't parse route")
	}

	link, err := su.repo.OneLink(userId, int(linkId))

	if err != nil {
		return nil, err
	}

	stat := Stat{
		Route: strconv.FormatInt(int64(link.Id), 36),
		Link:  link,
	}

	return &stat, nil
}

func (su *statUsecase) AllStat(userId int) ([]Stat, error) {
	links, err := su.repo.AllLinks(userId)

	if err != nil {
		return nil, err
	}

	stats := make([]Stat, 0, len(links))
	for i := range links {
		stats = append(stats, Stat{
			Route: strconv.FormatInt(int64(links[i].Id), 36),
			Link:  &links[i],
		})
	}
	return stats, nil
}

func newStatUsecase(repo repository.Statistics) *statUsecase {
	return &statUsecase{repo: repo}
}
