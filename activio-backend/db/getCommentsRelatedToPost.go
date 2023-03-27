package db

import (
	"log"

	"activio-backend/models"
)

func GetCommentsRelatedToPost(postID uint) ([]models.CommentApiResponse, error) {
	var comments []models.CommentApiResponse

	err := GetDB().Table("comments").Where("post_id = ?", postID).Find(&comments).Error

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return comments, nil
}