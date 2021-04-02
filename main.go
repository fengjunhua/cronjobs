package main

import (

   "github.com/gin-gonic/gin"
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

}

func main() {
   router := gin.Default()
   router.StaticFS("/static", http.Dir("./static"))
   router.LoadHTMLGlob("views/**/*")
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
   router.Run(":8090")
}