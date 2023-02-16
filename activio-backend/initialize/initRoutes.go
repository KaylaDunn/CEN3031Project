package initialize

import (
	"activio-backend/handlers/image"

	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api"


func InitRoutes(r *gin.Engine) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
  // as stated in the documentation:
  r.MaxMultipartMemory = 8 << 20  // 8 MiB

  // Alive endpoint
  r.GET(apiPrefix + "/alive", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "I'm alive!",
    })
  })

  // Single image endpoints
  r.Static(apiPrefix + "/image", "./images")
  r.POST(apiPrefix + "/image", image.UploadImage)

  // Multiple image endpoints
  r.POST(apiPrefix + "/images", image.MultipleImageUpload)
}