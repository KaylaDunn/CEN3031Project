package db

import (
	"log"

	"activio-backend/models"
)

func GetImagesRelatedToPost(postID uint) ([]models.Image, error) {
	var images []models.Image

	err := GetDB().Where("post_id = ?", postID).Find(&images).Error

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return images, nil
}