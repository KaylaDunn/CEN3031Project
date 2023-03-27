package handlers

import (
	"log"
	"net/http"

	"activio-backend/db"
	"activio-backend/models"

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
