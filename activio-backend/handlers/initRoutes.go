package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api"


func InitRoutes(r *gin.Engine) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
  // as stated in the documentation:
  r.MaxMultipartMemory = 8 << 20  // 8 MiB

  api := r.Group(apiPrefix)

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
  api.POST("/user", CreateNewUser)
}