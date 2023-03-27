package db

import (
	"log"

	"activio-backend/models"
)


func GetPosts(limit int, offset int) ([]models.Post, error) {
	var posts []models.Post
	db := GetDB()
	err := db.Statement.Order("created_at DESC").Limit(limit).Offset(offset).Find(&posts).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return posts, nil
}
