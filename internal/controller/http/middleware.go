package httpControl

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

func (c *Controller) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userId, err := c.usecases.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set("userId", userId)
}
