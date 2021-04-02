package controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context){

	c.HTML(200, "main.html", gin.H{
		"title": "ss",
	})
}

func Help(c *gin.Context)  {

	c.HTML(200, "help.html", gin.H{
		"title": "ss",
	})
}

func Start(c *gin.Context)  {
	c.HTML(200, "layout.html", gin.H{
		"title": "ss",
	})
}