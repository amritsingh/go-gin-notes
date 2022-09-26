package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("notes:tmp_pwd@tcp(127.0.0.1:3306)/notes?charset=utf8&parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Note{})
	DB.AutoMigrate(&User{})
}
