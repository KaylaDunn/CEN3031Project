package initialize

import (
	handlers "activio-backend/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api"


func InitRoutes(r *gin.Engine) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
  // as stated in the documentation:
  r.MaxMultipartMemory = 8 << 20  // 8 MiB

  // Alive endpoint
  r.GET(apiPrefix + "/alive", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "I'm alive!",
    })
  })

  // image endpoints
  r.Static(apiPrefix + "/image", "./images")
  r.POST(apiPrefix + "/image", handlers.UploadImage)
  r.POST(apiPrefix + "/images", handlers.MultipleImageUpload)

  // user endpoints
  r.POST(apiPrefix + "/user", handlers.CreateNewUser)
}