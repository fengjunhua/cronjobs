package models

import (
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init(){
    dns := "root:120225fjhFJH!@tcp(182.92.236.162:3306)/cronJobs?charset=utf8"
	db, err := gorm.Open("mysql",dns)
	defer db.Close()
    if err != nil{
    	panic(err)
	}
	log.Info("connect mysql success!")

}
