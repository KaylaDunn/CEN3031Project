package db

import (
	"activio-backend/models"
)

// GetPostsByLocation returns all posts by location
func GetPostsByLocation(location string) ([]models.PostApiResponse, error) {
	var posts []models.PostApiResponse

	res := GetDB().Table("posts").Where("location_name = ?", location).Find(&posts)

	if res.Error != nil {
		return nil, res.Error
	}
 
	return posts, nil
}