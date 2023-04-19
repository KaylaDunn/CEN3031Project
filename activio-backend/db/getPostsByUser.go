package db

import (
	"activio-backend/models"
)

// GetPostsByUser returns all posts by user
func GetPostsByUser(userID int) ([]models.PostApiResponse, error) {
	var posts []models.PostApiResponse

	res := GetDB().Table("posts").Where("user_id = ?", userID).Find(&posts)

	if res.Error != nil {
		return nil, res.Error
	}

	return posts, nil
}