package db

import (
	"activio-backend/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func InitDb() {
	dsn := "activio:activio@tcp(host.docker.internal:3306)/activio?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = db

	modelMigrate()
}

// Get db instance
func GetDB() *gorm.DB {
	return dbInstance
}

// Add models here
func modelMigrate() {
	dbInstance.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Image{})
}