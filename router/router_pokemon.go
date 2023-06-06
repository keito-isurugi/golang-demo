package router

import (
	"encoding/json"
	// "fmt"
	"golang_demo/db"
	"golang_demo/middleware"
	"log"
	"net/http"
	"strconv"
	"strings"
	// "net/url"
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
func getPokemonsPageHndler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	json.NewEncoder(w).Encode(db.GetPokemons(page))
}
func getPokemonListHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetPokemonList())
}


func RouterPokemon() {
	http.HandleFunc("/", homePage)

	// Pokemon
	http.HandleFunc("/api/pokemon/show/", middleware.CORS(getPokemonHndler))
	http.HandleFunc("/api/pokemon/list/", middleware.CORS(getPokemonsPageHndler))
	http.HandleFunc("/api/pokemons", middleware.CORS(getPokemonListHndler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
