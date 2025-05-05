package main

import (
	"go_task_api/controllers"
	"go_task_api/middlewares"
	"go_task_api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&models.User{}, &models.Task{})
}

func main() {
	initDatabase()

	// Inject DB into controllers
	controllers.InitAuth(DB)
	controllers.InitTask(DB)

	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected task routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetTasks)
		auth.POST("/tasks", controllers.CreateTask)
		auth.PUT("/tasks/:id", controllers.UpdateTask)
		auth.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	r.Run() // :8080
}
