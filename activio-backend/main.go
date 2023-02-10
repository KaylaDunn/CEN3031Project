package main

import (
	"github.com/gin-gonic/gin"

	image "activio-backend/handlers/image"
	utils "activio-backend/utils"
)

func main() {
  r := gin.Default()

  utils.CreateImageDir()

  // Set a lower memory limit for multipart forms (default is 32 MiB)
  // as stated in the documentation:
	r.MaxMultipartMemory = 8 << 20  // 8 MiB

  // Single image endpoints
  r.Static("/image", "./images")
  r.POST("/image", image.UploadImage)

  // Multiple image endpoints
  r.POST("/images", image.MultipleImageUpload)


  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}