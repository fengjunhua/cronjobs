package models

type Role struct {
	Id              int
	RoleName        string
	Detail          string
	ServerGroupIds  string
	TaskGroupsIds   string
	Status          int
	CreateId        int
	UpdateId        int
	CreateTime      int64
	UpdateTime      int64
}

func (a *Role) TableName() string {
	return "uc_role"
}

