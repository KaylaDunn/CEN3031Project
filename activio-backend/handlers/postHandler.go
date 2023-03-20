package handlers

import (
	"log"
	"net/http"

	"activio-backend/db"

	"github.com/gin-gonic/gin"
)

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

	// convert the posts to json
	// postsJSON, err := json.Marshal(posts)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	// return the posts
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
		"posts": posts,
	})
}