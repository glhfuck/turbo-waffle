package httpcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type shortControl struct {
	usecase usecase.Shortener
}

func newShortControl(uc usecase.Shortener) *shortControl {
	return &shortControl{usecase: uc}
}

func (sc *shortControl) redirect(ctx *gin.Context) {
	route := ctx.Param("route")

	originalURL, err := sc.usecase.ParseRoute(route)

	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, originalURL)
}

func (sc *shortControl) short(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	var input domain.Link
	err := ctx.BindJSON(&input)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	shortRoute, err := sc.usecase.ShortURL(input.OriginalURL, userId.(int))

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"URL": ctx.Request.URL.Host + "/" + shortRoute,
	})
}
