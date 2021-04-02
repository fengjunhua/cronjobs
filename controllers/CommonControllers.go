package controllers

import (
	"github/fengjunhua/cronjobs/models"
)

type BaseController struct {
	controllerName   string
	actionName       string
	user             *models.Admin
	userId           int
	userName         string
	loginName        string
	pageSize         int
	allowUrl         string
	serverGroups     string
	taskGroups       string
}

//前期准备
func (self *BaseController) Prepare(){
	self.pageSize = 20
	//controllerName,actionName := self.GetControllerAndAction()
	//self.controllerName = strings.ToLower(controllerName[0:len(controllerName)-10])
	//self.actionName = strings.ToLower(actionName)
}

//登录权限验证
func (self *BaseController) Auth()  {



}

func (self *BaseController) AdminAuth()  {
	//左侧导航栏
	filters := make([]interface{},0)
	filters = append(filters,"status",1)


}


















