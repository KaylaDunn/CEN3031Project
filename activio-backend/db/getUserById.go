package db

import (
	"activio-backend/models"

	"log"
)

// Get a User by its id
func GetUserById(id int) (models.UserApiResponse, error) {
	var user models.UserApiResponse
	err := GetDB().Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Println(err)
		return models.UserApiResponse{}, err
	}
	return user, nil
}