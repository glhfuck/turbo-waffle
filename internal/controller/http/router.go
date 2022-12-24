package httpControl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(c *Controller) http.Handler {
	engine := gin.New()

	auth := engine.Group("/auth")
	{
		auth.POST("/sign-up", c.signUp)
		auth.POST("/sign-in", c.signIn)
	}

	api := engine.Group("/", c.userIdentity)
	{
		api.GET(":id", c.redirect)

		api.POST("", c.short)

		stat := api.Group("stat")
		{
			stat.GET("/all", c.allStat)
			stat.GET("/:id", c.oneStat)
		}
	}

	return engine
}
