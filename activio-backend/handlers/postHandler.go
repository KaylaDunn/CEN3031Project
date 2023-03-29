package handlers

import (
	"log"
	"net/http"
	"strconv"

	"activio-backend/db"
	"activio-backend/models"
	"activio-backend/utils"

	"github.com/gin-gonic/gin"
)

type PostResponse struct {
	models.PostApiResponse
	Images []models.ImageApiResponse `json:"images"`
	Comments []models.CommentApiResponse `json:"comments"`
}

func GetPosts(c *gin.Context) {
	// get 10 posts from the database
	posts, err := db.GetPosts(10, 0)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	// create a PostResponse slice and fill it with the posts and their related data
	var postResponses []PostResponse

	for _, post := range posts {
		images, err := db.GetImagesRelatedToPost(post.ID)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		comments, err := db.GetCommentsRelatedToPost(post.ID)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		postResponses = append(postResponses, PostResponse{post, images, comments})
	}

	// return the posts
	c.JSON(http.StatusOK, gin.H{
		"posts": postResponses,
	})
}

func createPost(c *gin.Context) {
	// get the user id from the context
	userID := c.MustGet("user").(models.User).ID

	// get the post data from the request
	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in request body",
			"suggestion": "Check if the request body is valid JSON and if all required fields are present.",
		})
		return
	}

	// set the user id of the post
	post.UserID = userID

	// create the post in the database
	err = db.GetDB().Create(&post).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating post",
		})
		return
	}

	// return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPost(c *gin.Context) {
	// get the post id from the request
	postID := c.Param("id")

	// convert the post id to an int
	id, err := strconv.Atoi(postID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"suggestion": "Check if the ID is a number.",
		})
		return
	}

	// get the post from the database
	post, err := db.GetPostById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Post not found",
		})
		return
	}

	// get the images related to the post
	images, err := db.GetImagesRelatedToPost(post.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	// get the comments related to the post
	comments, err := db.GetCommentsRelatedToPost(post.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	// return the post
	c.JSON(http.StatusOK, gin.H{
		"post": PostResponse{post, images, comments},
	})
}

func AddImagesToPost (c *gin.Context) {
	// get the post id from the request
	postID := c.Param("id")

	// convert the post id to an int
	id, err := strconv.Atoi(postID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"suggestion": "Check if the ID is a number.",
		})
		return
	}

	// check to see if the post exists and if the user is the owner of the post
	post, err := db.GetPostById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post not found",
		})
		return
	}

	// get the user id from the context
	userID := c.MustGet("user").(models.User).ID

	// check to see if the user is the owner of the post
	if post.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to add images to this post",
		})
		return
	}

	// get the images from the request
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

	for i, file := range files {
		
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

		// Save the image to the database
		image := models.Image{
			UploadedBy: userID,
			PostID : post.ID,
			OriginalName : file.Filename,
			HashedFileName: compressedName,
			Order: i,
			IsProfilePicture: false,
		}

		err = db.GetDB().Create(&image).Error
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error creating post",
			})
			return
		}
	}

	// Return the list of filenames to the caller
	c.JSON(http.StatusOK, gin.H{"status": "ok", "filenames": filenames})
}