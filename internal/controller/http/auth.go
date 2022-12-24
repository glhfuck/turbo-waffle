package httpControl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
)

type authControl struct {
	usecases *usecase.Usecase
}

func newAuthControl(uc *usecase.Usecase) *authControl {
	return &authControl{usecases: uc}
}

func (ac *authControl) signUp(ctx *gin.Context) {
	var input domain.User

	err := ctx.BindJSON(&input)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ac.usecases.Authorization.CreateUser(input)

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

	token, err := ac.usecases.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
