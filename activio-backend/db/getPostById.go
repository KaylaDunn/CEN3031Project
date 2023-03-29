package db

import (
	"activio-backend/models"

	"log"
)

// Get a post by its id
func GetPostById(id int) (models.PostApiResponse, error) {
	var post models.PostApiResponse
	db := GetDB()
	err := db.Table("posts").Where("id = ?", id).First(&post).Error
	if err != nil {
		log.Println(err)
		return models.PostApiResponse{}, err
	}
	return post, nil
}