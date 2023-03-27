package db

import (
	"log"

	"activio-backend/models"
)

// Gets the limit most recent posts from the database with an offset of offset posts
func GetPosts(limit int, offset int) ([]models.PostApiResponse, error) {
	var posts []models.PostApiResponse
	db := GetDB()
	err := db.Table("posts").Order("created_at desc").Limit(limit).Offset(offset).Find(&posts).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return posts, nil
}
