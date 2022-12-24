package httpControl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) redirect(ctx *gin.Context) {
	id, _ := ctx.Get("userId")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (c *Controller) short(ctx *gin.Context) {

}
