package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(){
    dns := "root:120225fjhFJH!@tcp(182.92.236.162:3306)/cronJobs?charset=utf8"
	db, err := gorm.Open(mysql.Open(dns),&gorm.Config{})
    if err != nil{
    	panic(err)
	}
	fmt.Println("connect to mysql success!")
	fmt.Println(db)
}
