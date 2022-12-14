package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/rmediasolutions-website"))

	if err != nil {
		log.Println("connection failed", err)
	} else {
		log.Println("connection success")
	}

	db.AutoMigrate(&AboutUs{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&AdmUser{})
	DB = db
}
