// middleware/auth.go
package middlewares

import (
	"go_task_api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Package middleware handles JWT authentication.
// It protects routes by validating the Bearer token provided in the Authorization header.
//
// Example:
// Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR...
//
// If valid, it sets userID in the Gin context.
// If missing/invalid, it returns 401 Unauthorized.

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set userID in context
		c.Set("userID", claims.UserID)

		// Continue to next handler
		c.Next()
	}
}
