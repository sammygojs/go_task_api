package models

type Tag struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"urgent"`
}