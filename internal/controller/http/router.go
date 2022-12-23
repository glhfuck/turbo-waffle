package httpControl

import "github.com/gin-gonic/gin"

func NewRouter(e *gin.Engine, c Controller) {


	api := e.Group("/")
	{
		api.GET(":id", c.redirect)

		api.POST("", c.short)

		stat := api.Group("stat")
		{
			stat.GET("/all", c.allStat)
			stat.GET("/:id", c.oneStat)
		}

		auth := api.Group("auth")
		{
			auth.POST("/sign-up", c.signUp)
			auth.POST("/sign-in", c.signIn)
		}
	}
}
