package handlers

import (
	"activio-backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api"


func InitRoutes(r *gin.Engine) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
  // as stated in the documentation:
  r.MaxMultipartMemory = 8 << 20  // 8 MiB

  api := r.Group(apiPrefix) // General API group
  auth := api.Group("/auth", middleware.RequireAuth) // Auth group

  // Alive endpoint
  api.GET("/alive", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "I'm alive!",
    })
  })

  // image endpoints
  api.Static("/image", "./images")
  api.POST("/image", UploadImage)
  api.POST("/images", MultipleImageUpload)

  // user endpoints
  api.POST("/signup", SignUp)
  api.POST("/login", Login)
  auth.GET("/refreshtoken", RefreshJWT)
  auth.DELETE("/deleteaccount", DeleteUserAndUserData)
  auth.GET("/me", getUserDetails)

  // post endpoints
  api.GET("/posts", GetPosts)
  api.GET("/post/:id", GetPost)
  auth.POST("/createpost", createPost)
  auth.PUT("/addImageToPost/:id", AddImagesToPost)
}