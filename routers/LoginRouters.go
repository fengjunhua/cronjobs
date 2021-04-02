package routers

import (
	"github.com/gin-gonic/gin"
	"github/fengjunhua/cronjobs/controllers"
)

func LoadLoginRouters(e *gin.Engine) {

	LoginRouters := e.Group("/")
	{
		LoginRouters.GET("/", controllers.Login)
		LoginRouters.GET("/login", controllers.LoginIn)
		LoginRouters.GET("/logout", controllers.LoginOut)
		LoginRouters.GET("/reg", controllers.Register)
	}
}