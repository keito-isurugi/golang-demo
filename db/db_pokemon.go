package db

import (
	// "encoding/json"
	// "fmt"
	// "github.com/jinzhu/gorm"
	"golang_demo/models"
	_ "github.com/go-sql-driver/mysql"
)

// データ取得
func GetPokemons() []models.Pokemon {
	var pokemons []models.Pokemon
	ConnectDB()
	db.Find(&pokemons)
	defer db.Close()
	return pokemons
}
func GetPokemon(id int) models.Pokemon {
	var pokemon models.Pokemon
	ConnectDB()
	db.First(&pokemon, id)
	defer db.Close()
	return pokemon
}
