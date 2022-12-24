package httpControl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type shortControl struct {
	usecases *usecase.Usecase
}

func newShortControl(uc *usecase.Usecase) *shortControl {
	return &shortControl{usecases: uc}
}

func (sc *shortControl) redirect(ctx *gin.Context) {

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

	shortRoute, err := sc.usecases.Shortener.ShortURL(input.OriginalURL, userId.(int))

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"URL": ctx.Request.URL.Host + "/" + shortRoute,
	})
}
