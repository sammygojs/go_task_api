package controllers

import (
	"go_task_api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Admin-only: List all users
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.User
// @Router /admin/users [get]
func AdminGetUsers(c *gin.Context) {
	var users []models.User
	DB.Select("id", "username", "role", "created_at").Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Admin-only: Change user role (admin/user)
// @Tags Admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param role body map[string]string true "New role"
// @Success 200 {object} map[string]string
// @Failure 400,404 {object} map[string]string
// @Router /admin/users/{id}/role [put]
func AdminUpdateUserRole(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Role string `json:"role" example:"admin"` // "user" or "admin"
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if input.Role != "admin" && input.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role must be 'admin' or 'user'"})
		return
	}

	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Role = input.Role
	DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User role updated", "user": user.Username, "new_role": user.Role})
}
