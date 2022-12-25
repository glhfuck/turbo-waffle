package httpcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type Controller struct {
	Authorization
	Shortener
	Statistics
}

type Authorization interface {
	signUp(ctx *gin.Context)
	signIn(ctx *gin.Context)
	userIdentity(ctx *gin.Context)
}

type Shortener interface {
	redirect(ctx *gin.Context)
	short(ctx *gin.Context)
}

type Statistics interface {
	allStat(ctx *gin.Context)
	oneStat(ctx *gin.Context)
}

func NewController(uc *usecase.Usecase) *Controller {
	return &Controller{
		Authorization: newAuthControl(uc.Authorization),
		Shortener:     newShortControl(uc.Shortener),
		Statistics:    newStatControl(uc.Statistics),
	}
}
