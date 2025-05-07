package controllers

import (
	"go_task_api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var TaskDB *gorm.DB

func InitTask(db *gorm.DB) {
	TaskDB = db
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

// @Summary Get all tasks for the logged-in user (with filters, pagination, sorting)
// @Tags Tasks
// @Security BearerAuth
// @Produce json
// @Param project_id query int false "Filter by Project ID"
// @Param limit query int false "Max number of results"
// @Param offset query int false "Number of results to skip"
// @Param sort query string false "Sort by field (e.g. created_at, title, status)"
// @Param order query string false "Sort order (asc or desc)"
// @Success 200 {array} models.Task
// @Failure 401 {object} map[string]string
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	projectID := c.Query("project_id")
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")

	var tasks []models.Task
	query := TaskDB.Where("user_id = ?", userID)

	if projectID != "" {
		query = query.Where("project_id = ?", projectID)
	}

	if order != "asc" && order != "desc" {
		order = "desc" // fallback
	}

	query = query.Order(sort + " " + order).
		Limit(toInt(limit)).
		Offset(toInt(offset))

	query.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// @Summary Create a new task
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param task body models.Task true "Task info"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.UserID = userID
	TaskDB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// @Summary Update a task
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body models.Task true "Updated task"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var task models.Task
	if err := TaskDB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	TaskDB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// @Summary Delete a task
// @Tags Tasks
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 204
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var task models.Task
	if err := TaskDB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	TaskDB.Delete(&task)
	c.Status(http.StatusNoContent)
}
