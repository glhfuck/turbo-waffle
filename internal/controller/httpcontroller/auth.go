package httpcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type Authorization interface {
	signUp(ctx *gin.Context)
	signIn(ctx *gin.Context)
	userIdentity(ctx *gin.Context)
}

type authControl struct {
	usecase usecase.Authorization
}

func (ac *authControl) signUp(ctx *gin.Context) {
	var input domain.User

	err := ctx.BindJSON(&input)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ac.usecase.CreateUser(input)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

func (ac *authControl) signIn(ctx *gin.Context) {
	var input signInInput

	err := ctx.BindJSON(&input)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := ac.usecase.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func newAuthControl(uc usecase.Authorization) *authControl {
	return &authControl{usecase: uc}
}
