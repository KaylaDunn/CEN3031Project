package middleware

import (
	"activio-backend/db"
	"activio-backend/models"
	"activio-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)



func RequireAuth(c *gin.Context) {

	// get token from cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
	}

	// verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(utils.JWT_SECRET), nil
	})

	// Verify token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		}

		// Make sure user is still valid
		var user models.User
		db.GetDB().First(&user, claims["id"])

		if user.Email == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		}

		// Set user in context
		c.Set("user", user)


		// Continue
		c.Next()
	} else {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
	}
}