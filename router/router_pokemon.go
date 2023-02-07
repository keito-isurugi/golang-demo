package router

import (
	"encoding/json"
	// "fmt"
	"golang_demo/db"
	"log"
	"net/http"
	"strconv"
	"strings"
	"golang_demo/middleware"
	// "os"
	// _ "github.com/go-sql-driver/mysql"
	// gorm "github.com/jinzhu/gorm"
	// godotnev "github.com/joho/godotenv"
)


func getPokemonHndler(w http.ResponseWriter, r *http.Request) {
	todoId := strings.TrimPrefix(r.URL.Path, "/api/pokemon/show/")
	id, _ := strconv.Atoi(todoId)
	json.NewEncoder(w).Encode(db.GetPokemon(id))
}
func getPokemonsHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetPokemons())
}

func RouterPokemon() {
	http.HandleFunc("/", homePage)

	// Pokemon
	http.HandleFunc("/api/pokemon/show/", middleware.CORS(getPokemonHndler))
	http.HandleFunc("/api/pokemon/list", middleware.CORS(getPokemonsHndler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}