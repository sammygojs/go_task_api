package controllers

import (
	"go_task_api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var TagDB *gorm.DB

func InitTag(db *gorm.DB) {
	TagDB = db
}

// @Summary Create a new tag
// @Tags Tags
// @Accept json
// @Produce json
// @Param tag body models.Tag true "Tag info"
// @Success 201 {object} models.Tag
// @Router /tags [post]
func CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	TagDB.Create(&tag)
	c.JSON(http.StatusCreated, tag)
}

// @Summary Get all tags
// @Tags Tags
// @Produce json
// @Success 200 {array} models.Tag
// @Router /tags [get]
func GetTags(c *gin.Context) {
	var tags []models.Tag
	TagDB.Find(&tags)
	c.JSON(http.StatusOK, tags)
}
