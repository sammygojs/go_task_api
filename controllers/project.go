package controllers

import (
	"go_task_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ProjectDB *gorm.DB

func InitProject(db *gorm.DB) {
	ProjectDB = db
}

// @Summary Create a new project
// @Tags Projects
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param project body models.Project true "Project data"
// @Success 201 {object} models.Project
// @Router /projects [post]
func CreateProject(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project.UserID = userID
	ProjectDB.Create(&project)
	c.JSON(http.StatusCreated, project)
}

// @Summary Get all projects
// @Tags Projects
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Project
// @Router /projects [get]
func GetProjects(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var projects []models.Project
	ProjectDB.Where("user_id = ?", userID).Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// @Summary Get tasks for a specific project
// @Tags Projects
// @Security BearerAuth
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {array} models.Task
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /projects/{id}/tasks [get]
func GetProjectTasks(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	projectID := c.Param("id")

	// Optional: Check if project exists and belongs to user
	var project models.Project
	if err := ProjectDB.Where("id = ? AND user_id = ?", projectID, userID).First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var tasks []models.Task
	TaskDB.Where("project_id = ? AND user_id = ?", projectID, userID).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

