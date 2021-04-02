package config

type Application struct {
	AppName   string
	HttpPort  int
	RunMode   string
	Version   string
}


type Database struct {
	Host          string    `json:"host"`
	Port          string    `json:"port"`
	UserName      string    `json:"username"`
	Password      string    `json:"password"`
    Name          string    `json:"name"`
	Prefix        string    `json:"prefix"`
	Timezone      string    `json:"timezone"`
}

type Email struct {

}
type Msg struct {

}
type DingTalk struct {

}
type Wechat struct {

}
type Notify struct {
	 Email Email
     Msg  Msg
	 DingTalk DingTalk
	 Wechat  Wechat
}