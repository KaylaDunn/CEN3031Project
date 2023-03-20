package db

import (
	"log"

	"activio-backend/models"
)

func GetCommentsRelatedToPost(postID uint) ([]models.Comment, error) {
	var comments []models.Comment

	err := GetDB().Where("commented_on = ?", postID).Find(&comments).Error

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return comments, nil
}