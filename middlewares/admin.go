package middlewares

import (
	"go_task_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var AdminDB *gorm.DB

func InitAdmin(db *gorm.DB) {
	AdminDB = db
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(uint)

		var user models.User
		if err := AdminDB.First(&user, userID).Error; err != nil || user.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access only"})
			c.Abort()
			return
		}

		c.Next()
	}
}
