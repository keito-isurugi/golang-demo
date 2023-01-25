package models

import (
	"time"
)

type User struct {
	ID int `gorm:"primary_key" json:"id"`
  Name string `json:"name"`
	Email string `json:"email"`
	Password int `json:"password"`
	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
	Todos []Todo `json:"todos"`
}

type Todo struct {
	ID int `gorm:"primary_key" json:"id"`
	UserID int `json:"user_id"`
  Title string `json:"title"`
  Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
	User User `json:"user"`
}