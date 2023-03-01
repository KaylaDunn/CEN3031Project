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

  api := r.Group(apiPrefix, middleware.CORSMiddleware) // General API group
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
  api.OPTIONS("/*", func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, client-security-token")
    c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
    c.Writer.Header().Set("Access-Control-Max-Age", "86400")
    c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
      return
    }

  })
  api.POST("/login", Login)
  auth.GET("/refreshtoken", RefreshJWT)
}