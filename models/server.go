package models

import "fmt"

type TaskServer struct {
	Id              int
	GroupId         int
	ConnectionType  int
	ServerName      string
	ServerAccount   string
	ServerOuterIp   string
	ServerIp        string
	Port            int
	Password        string
	PrivateKeySrc   string
	PublicKeySrc    string
	Type            int
	Detail          string
	CreateTime      int64
	UpdateTime      int64
	Status          int
}

func (t *TaskServer) TableName() string {
	return "task_server"
}

func (t *TaskServer) Updata(fields ...string) error{
	if t.ServerName == "" {
		return fmt.Errorf("服务器名不能为空")
	}
	if t.ServerIp == "" {
		return fmt.Errorf("服务器IP不能为空")
	}
	if t.ServerAccount == "" {
		return fmt.Errorf("登录账户不能为空")
	}
	if t.Type == 0 && t.Password == "" {
		return fmt.Errorf("服务器密码不能为空")
	}
	if t.Type == 1 && t.PrivateKeySrc == "" {
		return fmt.Errorf("私钥不能为空")
	}
	/**

	更新服务器数据库操作
	 */

	return nil
}

func TaskServerAdd(obj *TaskServer) (int64,error) {
	if obj.ServerName == "" {
		return 0,fmt.Errorf("服务器名不能为空")
	}
	if obj.ServerIp == "" {
		return 0,fmt.Errorf("服务器IP不能为空")
	}
	if obj.ServerAccount == "" {
		return 0,fmt.Errorf("登录账户不能为空")
	}
	if obj.Type == 0 && obj.Password == "" {
		return 0,fmt.Errorf("服务器密码不能为空")
	}
	if obj.Type == 1 && obj.PrivateKeySrc == "" {
		return 0,fmt.Errorf("私钥不能为空")
	}
	/*
	创建数据库服务器操作
	 */
	return 0, nil
}

func GetTaskServerById(id int)(*TaskServer,error){
	obj := &TaskServer{
		Id: id,
	}
	/*
	获取服务器
	 */

	return obj,nil
}


