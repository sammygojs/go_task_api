package models

type Project struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"Work"`
	UserID uint   `json:"user_id" example:"2"`
}