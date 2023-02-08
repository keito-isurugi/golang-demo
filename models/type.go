package models

import (
	"encoding/json"
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
	Hp             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"special_attack"`
	SpecialDefense int `json:"special_defense"`
	Speed          int `json:"speed"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type PokemonStatus struct {
	Hp             json.Number `json:"hp"`
	Attack         json.Number `json:"attack"`
	Defense        json.Number `json:"defense"`
	SpecialAttack  json.Number `json:"special_attack"`
	SpecialDefense json.Number `json:"special_defense"`
	Speed          json.Number `json:"speed"`
}
