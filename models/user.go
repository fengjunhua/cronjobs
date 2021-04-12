package models

type User struct {
	Id          int
	LoginName   string
	RealName    string
	Password    string
	RoleIds     string
	Phone       string
	Email       string
	DingTalk    string
	Wechat      string
	Salt        string
	LastLogin   int64
	LastIp      string
	Status      int
	CreateId    int
	UpdateId    int
	CreateTime  int64
	UpdateTime  int64
}
