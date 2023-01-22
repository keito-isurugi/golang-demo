package models

import (
	"time"
)

type User struct {
	Id int `gorm:"primary_key"`
  Name string
	CreatedAt time.Time
  UpdatedAt time.Time
}

type Todo struct {
	Id int `gorm:"primary_key"`
  Title  string
  Content  string
	CreatedAt time.Time
  UpdatedAt time.Time
}