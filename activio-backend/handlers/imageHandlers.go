package handlers

import (
	"log"
	"net/http"

	"activio-backend/db"
	"activio-backend/models"
	"activio-backend/utils"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	// This endpoint is used to upload an image to the server
	// This will be used to upload profile pictures

	// Get userID from the context
	user, ok := c.Get("user")
	if !ok {
		log.Println("Error getting user from context")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid User Request", "Suggestion": "Please check the documentation for the correct request format"})
		return
	}


	file, err := c.FormFile("image")

	// Check if there is an error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error uploading file"})
		return
	}

	// Save image to server
	hashedFilename, err := utils.SaveImage(file, 50)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error uploading file"})
		return
	}

	// Print the file name and time of upload
	log.Printf("File: %s uploaded\n", file.Filename)

	// Save the file name to the database
	db.SetProfilePicture(hashedFilename, user.(models.User).ID)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "hash name": hashedFilename})
}

func MultipleImageUpload(c *gin.Context) {
	// This endpoint is used to upload multiple images to the server

	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error uploading file"})
		return
	}

	files := form.File["images"]

	// Make sure there are files to upload
	if len(files) == 0 {
		log.Println("No files to upload")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No files to upload"})
		return
	}

	// Store all uploaded filenames to return to the caller
	filenames := []string{}

	for _, file := range files {
		
		// Save image to server
		compressedName, err := utils.SaveImage(file, 50)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Error uploading file"})
			return
		}

		// Append the filename to the list of filenames
		filenames = append(filenames, compressedName)
		
		// Print the file name and time of upload
		log.Printf("File: %s uploaded\n", file.Filename) 
	}

	// Return the list of filenames to the caller
	c.JSON(http.StatusOK, gin.H{"status": "ok", "filenames": filenames})
}