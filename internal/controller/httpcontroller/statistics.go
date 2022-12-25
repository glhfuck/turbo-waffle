package httpcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type statControl struct {
	usecase usecase.Statistics
}

type Stat struct {
	Route string `json:"route"`
	Link  *domain.Link
}

func newStatControl(uc usecase.Statistics) *statControl {
	return &statControl{usecase: uc}
}

func (sc *statControl) allStat(ctx *gin.Context) {

}

func (sc *statControl) oneStat(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	route := ctx.Param("route")

	link, err := sc.usecase.OneStat(userId.(int), route)

	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	stat := Stat{
		Route: strconv.FormatInt(int64(link.Id), 36),
		Link:  link,
	}

	ctx.JSON(http.StatusOK, stat)
}
