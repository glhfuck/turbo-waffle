package httpcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type Statistics interface {
	allStat(ctx *gin.Context)
	oneStat(ctx *gin.Context)
}

type statControl struct {
	usecase usecase.Statistics
}

func (sc *statControl) allStat(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	stats, err := sc.usecase.AllStat(userId.(int))
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, stats)
}

func (sc *statControl) oneStat(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	route := ctx.Param("route")
	stat, err := sc.usecase.OneStat(userId.(int), route)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, stat)
}

func newStatControl(uc usecase.Statistics) *statControl {
	return &statControl{usecase: uc}
}
