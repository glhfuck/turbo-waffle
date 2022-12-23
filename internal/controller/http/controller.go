package httpControl

import (
	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type Controller struct {
	Authorization
	Shortener
	Statistics

	usecases *usecase.Usecase
}

type Authorization interface {
	signUp(ctx *gin.Context)
	signIn(ctx *gin.Context)
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
	return &Controller{usecases: uc}
}
