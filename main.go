package main

import (
   "github.com/gin-gonic/gin"
   log "github.com/sirupsen/logrus"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
   _ "github/fengjunhua/cronjobs/docs"
   "github/fengjunhua/cronjobs/models"
   "github/fengjunhua/cronjobs/routers"
   "net/http"
)

var secrets = gin.H{
   "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
   "austin": gin.H{"email": "austin@example.com", "phone": "666"},
   "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func init()  {

   models.Init()
   log.SetFormatter(&log.JSONFormatter{})
   log.Println()
}

// @title 定时任务管理系统
// @version 1.0
// @description 定时任务管理系统后端API接口文档
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8090
// @BasePath /api/v1
func main() {
   router := gin.Default()
   router.StaticFS("/static", http.Dir("./static"))
   router.LoadHTMLGlob("views/**/*")
   url := ginSwagger.URL("http://localhost:8090/swagger/doc.json")
   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,url))
   authorized := router.Group("/admin",gin.BasicAuth(gin.Accounts{
      "foo":     "bar",
      "austin":  "1234",
      "lena":    "hello2",
      "manu":     "4321",
   }))
   authorized.GET("/secrets", func(c *gin.Context) {
      user := c.MustGet(gin.AuthUserKey).(string)
      if secret,ok := secrets[user];ok{
         c.JSON(http.StatusOK,gin.H{"user":user,"secret":secret})
      }else {
         c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
      }

   })
   routers.LoadLoginRouters(router)
   routers.LoadHomeRouters(router)
   routers.LoadApiRouters(router)
   router.Run(":8090")
}