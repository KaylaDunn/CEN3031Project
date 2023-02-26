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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
		return
	}

	// Check to see if the user already exists
	var existingUser models.User
	db.GetDB().Find(&existingUser, "email = ?", user.Email)

	if existingUser.Email != "" {
		log.Println("User already exists")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error logging in"})
		return
	}

	// Check to see if the user exists
	var existingUser models.User
	db.GetDB().Find(&existingUser, "email = ?", user.Email)

	if existingUser.Email == "" {
		log.Println("User does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	// Check to see if the password is correct
	if !utils.ComparePassword(user.Password, existingUser.Password) {
		log.Println("Incorrect password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	// Generate a jwt token
	token, err := utils.GenerateToken(existingUser.ID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error logging in"})
		return
	}

	// Return the jwt as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}