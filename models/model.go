package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pascal_practice_server/pkg/setting"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)), &gorm.Config{})

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Tag{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
