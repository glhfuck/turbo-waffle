package httpcontroller

import (
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type Controller struct {
	Authorization
	Shortener
	Statistics
}

func NewController(uc *usecase.Usecase) *Controller {
	return &Controller{
		Authorization: newAuthControl(uc.Authorization),
		Shortener:     newShortControl(uc.Shortener),
		Statistics:    newStatControl(uc.Statistics),
	}
}
