package db

import (
	"activio-backend/models"
	"log"
)

// GetPostsByLocation returns all posts by location
func GetPostsByLocation(location string) ([]models.PostApiResponse, error) {
	var posts []models.PostApiResponse

	log.Println("Location: ", location)

	res := GetDB().Table("posts").Where("location_name = ?", location).Find(&posts)

	if res.Error != nil {
		return nil, res.Error
	}

	log.Println("Posts by location: ", posts)
 
	return posts, nil
}