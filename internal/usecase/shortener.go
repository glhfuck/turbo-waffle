package usecase

import (
	"strconv"
	"time"

	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
)

const (
// salt       = "jHFGYkV7TvruGlni7ynIr5"
// signingKey = "JKv89Hfdf98ffSD"
// tokenTTL   = 12 * time.Hour
)

type shortUsecase struct {
	repo repository.Shortener
}

func newShortUsecase(repo repository.Shortener) *shortUsecase {
	return &shortUsecase{repo: repo}
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
