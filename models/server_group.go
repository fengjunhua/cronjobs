package models

import "fmt"

type ServerGroup struct {
	Id           int
	CreateId     int
	UpdateId     int
	GroupName    string
	Description  string
	CreateTime   int64
	UpdateTime   int64
	Status       int
}

func (t *ServerGroup) TableName() string {
	return "task_server_group"
}

func (t *ServerGroup) Update(field ...string) error {
	if t.GroupName == ""{
		return fmt.Errorf("组名不能为空")
	}
	/*

	 */

	return nil
}

func GetTaskGroupById(id int)(*ServerGroup,error){
	obj := &ServerGroup{
		Id: id,
	}
	/*

	 */

	return obj,nil
}