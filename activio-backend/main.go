package main

import (
	"github.com/gin-gonic/gin"

	db "activio-backend/db"
	handlers "activio-backend/handlers"
	utils "activio-backend/utils"
)

func main() {
  r := gin.Default()
  
  db.InitDb()
  handlers.InitRoutes(r)

  utils.CreateImageDir()


  r.Run() // listen and serve on 0.0.0.0:8080
}