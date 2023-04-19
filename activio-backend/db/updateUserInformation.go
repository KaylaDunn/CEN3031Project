package db

import (
	"log"

	"activio-backend/models"
)

// There's probably a better way to do this

func UpdateUser(user models.UserUpdateInfo, userID uint) error {

	// Find the user
	var userToUpdate models.User
	GetDB().First(&userToUpdate, userID)

	if user.FirstName != "" {
		userToUpdate.FirstName = user.FirstName
	}

	if user.LastName != "" {
		userToUpdate.LastName = user.LastName
	}

	if user.Username != "" {
		userToUpdate.Username = user.Username
	}

	if user.Phonenumber != "" {
		userToUpdate.Phonenumber = user.Phonenumber
	}

	if user.Birthday != "" {
		userToUpdate.Birthday = user.Birthday
	}

	// Save the user
	err := GetDB().Save(&userToUpdate).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
