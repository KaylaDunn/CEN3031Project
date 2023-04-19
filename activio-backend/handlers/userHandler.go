package handlers

import (
	"log"
	"net/http"

	"activio-backend/db"
	"activio-backend/models"
	"activio-backend/utils"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	// This endpoint is used to create a new user

	// Create a new user object
	var user models.User

	// Bind the request body to the user object
	err := c.BindJSON(&user)

	// Check if there is an error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User Request", "Suggestion": "Please check the documentation for the correct request format"})
		return
	}

	// Check to see if the user already exists
	var existingUser models.User
	db.GetDB().Find(&existingUser, "email = ?", user.Email)

	if existingUser.Email != "" {
		log.Println("User already exists")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User already exists"})
		return
	}

	// set the user to not verified
	user.Verified = false

	// hash the password
	user.Password = utils.HashPassword(user.Password)

	// Create the user
	db.GetDB().Create(&user)

	// Return the user's email, id, verified status
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"verified": user.Verified,
		"username": user.Username,
		"firstName": user.FirstName,
		"lastName": user.LastName,
		"birthday": user.Birthday,
		"phoneNumber": user.Phonenumber,
		"profilePicture": user.ProfilePicture,
	})
}

func Login(c *gin.Context) {
	// This endpoint is used to login a user

	// Create a new user object
	var user models.User

	// Bind the request body to the user object
	err := c.BindJSON(&user)

	// Check if there is an error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User Request", "Suggestion": "Please check the documentation for the correct request format"})
		return
	}

	// Check to see if the user exists
	var existingUser models.User
	db.GetDB().Find(&existingUser, "email = ?", user.Email)

	if existingUser.Email == "" {
		log.Println("User does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User does not exist", "Suggestion": "Please sign up or check your email"})
		return
	}

	// Check to see if the password is correct
	if !utils.ComparePassword(user.Password, existingUser.Password) {
		log.Println("Incorrect password and/or email")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Incorrect password and/or email"})
		return
	}

	// Generate a jwt token
	token, err := utils.GenerateToken(existingUser.ID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error logging in", "Suggestion": "Please try again later"})
		return
	}

	// Return the jwt as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func RefreshJWT(c *gin.Context) {
	// This endpoint is used to refresh the jwt token

	// Get the user from the context
	user, ok := c.Get("user")

	// Check if there is an error
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error refreshing token", "Suggestion": "Please try again later"})
		return
	}


	// Generate a jwt token
	token, err := utils.GenerateToken(user.(models.User).ID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error refreshing token", "Suggestion": "Please try again later"})
		return
	}

	// Return the jwt as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteUserAndUserData(c *gin.Context) {
	// This endpoint is used to delete a user and all of their data

	// Get the user from the context
	user, ok := c.Get("user")

	// Check if there is an error
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error deleting user", "Suggestion": "Please try again later"})
		return
	}

	// Verify if the user still exists
	var existingUser models.User
	db.GetDB().Find(&existingUser, "id = ?", user.(models.User).ID)

	if existingUser.Email == "" {
		log.Println("User does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User does not exist", "Suggestion": "Please sign up or check your email"})
		return
	}

	// Delete the user
	db.GetDB().Delete(&models.User{}, user.(models.User).ID)

	// TODO: Delete all of the user's data
	
	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func getUserDetails(c *gin.Context) {
	// Return user details

	// Get the user from the context
	user, ok := c.Get("user")

	// Check if there is an error
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error getting user details", "Suggestion": "Please try again later"})
		return
	}

	// Return the user's email, id, verified status
	c.JSON(http.StatusOK, gin.H{
		"id":       user.(models.User).ID,
		"email":    user.(models.User).Email,
		"verified": user.(models.User).Verified,
		"firstname": user.(models.User).FirstName,
		"lastname": user.(models.User).LastName,
		"phoneNumber": user.(models.User).Phonenumber,
		"username": user.(models.User).Username,
		"birthday": user.(models.User).Birthday,
	})
}

func UpdateUser(c *gin.Context) {
	// This endpoint is used to update a user's details

	// Get the user from the context
	user, ok := c.Get("user")

	// Check if there is an error
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error updating user", "Suggestion": "Please try again later"})
		return
	}

	// Verify if the user still exists
	var existingUser models.User
	db.GetDB().Find(&existingUser, "id = ?", user.(models.User).ID)

	if existingUser.Email == "" {
		log.Println("User does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User does not exist", "Suggestion": "Please sign up or check your email"})
		return
	}

	// Create a new user object
	var updatedUser models.UserUpdateInfo

	// Bind the request body to the user object
	err := c.BindJSON(&updatedUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User Request", "Suggestion": "Please check the documentation for the correct request format"})
		return
	}

	// print the user object
	log.Println(updatedUser.FirstName)
	log.Println(updatedUser.LastName)
	log.Println(updatedUser.Username)
	log.Println(updatedUser.Birthday)
	log.Println(updatedUser.Phonenumber)

	// Update the user
	err = db.UpdateUser(updatedUser, user.(models.User).ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error updating user", "Suggestion": "Please try again later"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}