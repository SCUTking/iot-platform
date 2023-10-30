package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:852456jkl@tcp(127.0.0.1:3306)/iot_platform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("[DB error]", err)
	}

	//数据库迁移——就是利用model文件构建sql表的

	err = db.AutoMigrate(&UserBasic{}, &DeviceBasic{}, &ProductBasic{})
	if err != nil {
		log.Fatalln("[DB error]", err)
	}
	DB = db
}
