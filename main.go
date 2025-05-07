package main

// @title Task Manager API
// @version 1.0
// @description This is a task manager backend built with Go.
// @host localhost:8080
// @BasePath 
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
import (
	ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
	"go_task_api/controllers"
	"go_task_api/middlewares"
	"go_task_api/models"
	_ "go_task_api/docs" 
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
	// @Summary Register a new user
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param input body map[string]string true "User credentials"
	// @Success 201 {object} map[string]string
	// @Router /register [post]
	r.POST("/register", controllers.Register)

	// @Summary Login using credentials
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param input body map[string]string true "User credentials"
	// @Success 200 {object} map[string]string
	// @Router /login [post]
	r.POST("/login", controllers.Login)

	//swagger routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))

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
