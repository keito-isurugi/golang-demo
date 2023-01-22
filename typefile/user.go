package typefile

import (
	// "github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id int `gorm:"primary_key"`
  Name string
	CreatedAt time.Time
  UpdatedAt time.Time
}

type Todo struct {
	id int `gorm:"primary_key"`
  Title  string
  Content  string
	CreatedAt time.Time
  UpdatedAt time.Time
}