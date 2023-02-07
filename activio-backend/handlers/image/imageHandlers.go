package image

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// file path to the image directory
const imageDir = "images/"

func GetImage(c *gin.Context) {
	log.Println("Get Image")
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func UploadImage(c *gin.Context) {
	// This endpoint is used to upload an image to the server

	file, err := c.FormFile("image")

	// Check if there is an error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
		return
	}

	// Check to see if the file is an image (jpg, png, etc.)
	fileType := file.Header["Content-Type"][0]
	if fileType != "image/jpeg" && fileType != "image/png" {
		log.Println("File is not an image")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
		return
	}

	// Save the file to the server
	// TODO: Hash the file name for storage
	// TODO: Compress the image
	err = c.SaveUploadedFile(file, imageDir + file.Filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not uploaded"})
		return
	}

	// Print the file name and time of upload
	log.Printf("File: %s uploaded\n", file.Filename)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}