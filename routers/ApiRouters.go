package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/cronjobs/controllers"
)


func LoadApiRouters(e *gin.Engine) {

	ApiRouters := e.Group("/api/v1")
	{

		ApiRouters.GET("/get", controllers.Get)
		ApiRouters.GET("/help", controllers.Help)
		ApiRouters.GET("/start", controllers.Start)

	}
}