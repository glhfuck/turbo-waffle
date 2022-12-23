package httpControl

import "github.com/gin-gonic/gin"

type Controller interface {
	redirect(ctx *gin.Context)
	short(ctx *gin.Context)

	allStat(ctx *gin.Context)
	oneStat(ctx *gin.Context)

	signUp(ctx *gin.Context)
	signIn(ctx *gin.Context)
}

type controller struct {

}

func NewController() Controller {
	return &controller{}
}


