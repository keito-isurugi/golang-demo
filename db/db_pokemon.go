package db

import (
	// "encoding/json"
	// "fmt"
	// "github.com/jinzhu/gorm"
	"golang_demo/models"
	_ "github.com/go-sql-driver/mysql"
)

// データ取得(ページネーション用)
func GetPokemons(page int) []models.Pokemon {
	numPerPage := 49
	limit := numPerPage
	offset := (page - 1) * numPerPage
	var pokemons []models.Pokemon
	ConnectDB()
	db.Limit(limit).Offset(offset).Find(&pokemons)
	defer db.Close()
	return pokemons
}
func GetPokemonList() []models.Pokemon {
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
