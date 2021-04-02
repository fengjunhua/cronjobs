package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/cronjobs/controllers"
)

func LoadHomeRouters(e *gin.Engine) {

	HomeRouters := e.Group("/home")
	{
		HomeRouters.GET("/", controllers.Index)
		HomeRouters.GET("/help", controllers.Help)
		HomeRouters.GET("/start", controllers.Start)

	}
}
