package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"go_task_api/models"
	// "go_task_api/utils"
	"go_task_api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTaskTestEnv() *gin.Engine {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Task{})
	InitAuth(db)
	InitTask(db)

	r := gin.Default()

	// Auth routes
	r.POST("/register", Register)
	r.POST("/login", Login)

	// Task routes (protected)
	taskGroup := r.Group("/tasks")
	taskGroup.Use(middleware.AuthMiddleware())
	{
		taskGroup.GET("", GetTasks)
		taskGroup.POST("", CreateTask)
	}

	return r
}

func registerAndLogin(r *gin.Engine, t *testing.T) string {
	// Register user
	regPayload := `{"username": "sumit", "password": "secret"}`
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(regPayload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("Register failed: %d", w.Code)
	}

	// Login user
	loginPayload := `{"username": "sumit", "password": "secret"}`
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/login", bytes.NewBufferString(loginPayload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Login failed: %d", w.Code)
	}
	var resp map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	return resp["token"]
}

func TestCreateAndGetTasks(t *testing.T) {
	r := setupTaskTestEnv()
	token := registerAndLogin(r, t)

	// Create a task
	taskPayload := `{"title": "Test Task", "status": "in-progress"}`
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(taskPayload))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", w.Code)
	}

	// Get tasks
	req = httptest.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}

	var tasks []models.Task
	if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}
	if len(tasks) != 1 || tasks[0].Title != "Test Task" {
		t.Fatalf("Unexpected tasks response: %+v", tasks)
	}
}
