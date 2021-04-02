package controllers

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "ss",
	})
}

func LoginIn(c *gin.Context) {


}

func LoginOut(c *gin.Context) {

}

func Register(c *gin.Context) {

}