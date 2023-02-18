package main

import (
	"github.com/gin-gonic/gin"

	initialize "activio-backend/initialize"
	utils "activio-backend/utils"
)

func main() {
  r := gin.Default()
  
  initialize.InitDb()
  initialize.InitRoutes(r)

  utils.CreateImageDir()


  r.Run() // listen and serve on 0.0.0.0:8080
}