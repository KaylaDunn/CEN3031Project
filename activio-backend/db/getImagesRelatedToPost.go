package db

import (
	"log"

	"activio-backend/models"
)

func GetImagesRelatedToPost(postID uint) ([]models.ImageApiResponse, error) {
	var images []models.ImageApiResponse

	err := GetDB().Table("images").Where("post_id = ?", postID).Find(&images).Error

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return images, nil
}