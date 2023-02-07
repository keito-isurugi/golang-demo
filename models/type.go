package models

import (
	"time"
)

type User struct {
	ID int `gorm:"primary_key" json:"id"`
  Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
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

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type1 string `json:"type1"`
	Type2 string `json:"type2"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Flavor_text string `json:"flavor_text"`
	Img string `json:"img"`
	Generation int `json:"generation"`
	Classification string `json:"classification"`
	Status string `json:"status"`
	// Status map[string]interface{} `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}