package main

import (
	"github.com/gin-gonic/gin"

	initialize "activio-backend/initialize"
	utils "activio-backend/utils"
)

func main() {
  r := gin.Default()

  utils.CreateImageDir()

  initialize.InitRoutes(r)


  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}