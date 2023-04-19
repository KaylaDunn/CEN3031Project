package db

import (
	"log"

	"activio-backend/models"
)

func SetProfilePicture(fileName string, userID uint) {
	// This function is used to set the profile picture of a user
	// It takes in a file name and a user id

	// Get the database connection
	db := GetDB()

	// Create a new user
	var user models.User

	// Get the user from the database
	db.First(&user, userID)

	// Set the profile picture
	user.ProfilePicture = fileName

	// Save the user to the database
	db.Save(&user)

	log.Printf("Profile picture for user %d set to %s\n", userID, fileName)
}