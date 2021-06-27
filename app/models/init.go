package models

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/27 10:11

import (
	"etl-demo/app/utils"
	"fmt"
	//"github.com/google/martian/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUp(isOrmDebug bool) *gorm.DB {
	config := utils.GetConfig()
	conUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.Db.User,
		config.Db.Passwd,
		config.Db.Host,
		config.Db.Port,
		config.Db.Db,
		config.Db.Charset)
	db, err := gorm.Open(mysql.Open(conUri), &gorm.Config{})
	if err != nil {
		panic("connect DB failed")
	}
	db.AutoMigrate(&UpLoadFile{}, &ExtractFile{})
	DB = db
	return db
}
