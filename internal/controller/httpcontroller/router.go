package httpcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(c *Controller) http.Handler {
	engine := gin.New()

	engine.GET("/:route", c.redirect)

	auth := engine.Group("/auth")
	{
		auth.POST("/sign-up", c.signUp)
		auth.POST("/sign-in", c.signIn)
	}

	api := engine.Group("/", c.userIdentity)
	{
		api.POST("short", c.short)

		stat := api.Group("stat")
		{
			stat.GET("/all", c.allStat)
			stat.GET("/:route", c.oneStat)
		}
	}

	return engine
}
