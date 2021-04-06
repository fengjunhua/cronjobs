package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// @Tags 用户信息
// @Summary 获取字符串
// @Description 接口详细描述信息
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /get [get]
func Get(c *gin.Context) {

	c.JSON(http.StatusOK,gin.H{"data":"ok"})

}