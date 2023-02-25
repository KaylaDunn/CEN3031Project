package db

import (
	"activio-backend/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	dsn := "activio:activio@tcp(host.docker.internal:3306)/activio?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	modelMigrate()
}

// Add models here
func modelMigrate() {
	DB.AutoMigrate(&models.User{})
}