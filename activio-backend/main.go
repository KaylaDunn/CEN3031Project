package main

import (
	"github.com/gin-gonic/gin"

	image "activio-backend/handlers/image"
)

func main() {
  r := gin.Default()

  // Set a lower memory limit for multipart forms (default is 32 MiB)
  // as stated in the documentation:
	r.MaxMultipartMemory = 8 << 20  // 8 MiB

  r.GET("/image", image.GetImage)
  r.POST("/image", image.UploadImage)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}