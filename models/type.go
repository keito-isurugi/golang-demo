package models

import (
	"time"
)

type User struct {
	ID int `gorm:"primary_key"`
  Name string
	Email string
	Password int
	CreatedAt time.Time
  UpdatedAt time.Time
	Todos []Todo
}

type Todo struct {
	ID int `gorm:"primary_key"`
	UserID int
  Title string
  Content string
	CreatedAt time.Time
  UpdatedAt time.Time
	User User
}