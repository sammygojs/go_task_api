package models

// RegisterRequest represents the payload for user registration
type RegisterRequest struct {
	Username string `json:"username" example:"sumit"`
	Password string `json:"password" example:"password123"`
}

// LoginRequest represents the payload for user login
type LoginRequest struct {
	Username string `json:"username" example:"sumit"`
	Password string `json:"password" example:"password123"`
}
