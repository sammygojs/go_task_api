package models

type Task struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"Buy milk"`
	Status    string `json:"status" example:"in-progress"`
	UserID    uint   `json:"user_id" example:"2"`
	CreatedAt string `json:"created_at" example:"2025-05-07T12:34:56Z"`
	UpdatedAt string `json:"updated_at" example:"2025-05-07T13:34:56Z"`
}
