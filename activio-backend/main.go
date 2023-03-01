package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"activio-backend/db"
	"activio-backend/handlers"
	"activio-backend/utils"
)

func main() {
  r := gin.Default()

  // Allow cors on localhost:3000, localhost:4200 for development
  r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:3000", "http://localhost:4200"},
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
    AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
    ExposeHeaders: []string{"Content-Length", "Content-Type", "Authorization"},
    AllowCredentials: true,
  }))
  
  db.InitDb()
  handlers.InitRoutes(r)

  utils.CreateImageDir()


  r.Run() // listen and serve on 0.0.0.0:8080
}