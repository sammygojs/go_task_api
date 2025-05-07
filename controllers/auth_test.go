// controllers/auth_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"go_task_api/models"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Task{})
	InitAuth(db)
	return db
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", Register)
	r.POST("/login", Login)
	return r
}

func TestRegisterAndLogin(t *testing.T) {
	db := setupTestDB()
	_ = db // to avoid lint error for unused variable
	r := setupRouter()

	// Step 1: Register
	registerPayload := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ := json.Marshal(registerPayload)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201, got %d", w.Code)
	}

	// Step 2: Login
	loginPayload := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ = json.Marshal(loginPayload)
	req = httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", w.Code)
	}

	var resp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Invalid JSON response")
	}
	if resp["token"] == "" {
		t.Fatalf("Token not returned")
	}
}
