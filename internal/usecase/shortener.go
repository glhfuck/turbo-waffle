package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
)

type shortUsecase struct {
	repo repository.Shortener
}

func newShortUsecase(repo repository.Shortener) *shortUsecase {
	return &shortUsecase{repo: repo}
}

func (su *shortUsecase) ParseRoute(route string) (string, error) {
	linkId, err := strconv.ParseInt(route, 36, 32)

	if err != nil {
		return "", errors.New("can't parse route")
	}

	originalURL, err := su.repo.OriginalURL(int(linkId))

	if err != nil {
		fmt.Println(err)
		return "", errors.New("no such route")
	}

	return originalURL, nil
}

func (su *shortUsecase) ShortURL(originalURL string, userId int) (string, error) {
	time := time.Now()

	link := &domain.Link{
		OwnerId:      userId,
		OriginalURL:  originalURL,
		CreationDate: time,
		UpdateDate:   time,
		VisitsCount:  0,
	}

	linkWithId, err := su.repo.SaveLink(link)

	if err != nil {
		return "", err
	}

	id := linkWithId.Id

	str := strconv.FormatInt(int64(id), 36)

	return str, nil
}
