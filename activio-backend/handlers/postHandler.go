package handlers

import (
	"log"
	"net/http"

	"activio-backend/db"
	"activio-backend/models"

	"github.com/gin-gonic/gin"
)

type PostResponse struct {
	models.Post
	Images []models.Image `json:"images"`
	Comments []models.Comment `json:"comments"`
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