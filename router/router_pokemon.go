package router

import (
	"encoding/json"
	"fmt"
	"golang_demo/db"
	"golang_demo/middleware"
	"log"
	"net/http"
	"strconv"
	"strings"
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

type Status struct {
	Hp             json.Number `json:"hp"`
	Attack         json.Number `json:"attack"`
	Defense        json.Number `json:"defense"`
	SpecialAttack  json.Number `json:"special_attack"`
	SpecialDefense json.Number `json:"special_defense"`
	Speed          json.Number `json:"speed"`
}

func getJsonDemoHndler(w http.ResponseWriter, r *http.Request) {

	pokestatus := `{"hp": 100, "attack": 100, "defense": 100, "special_attack": 100, "special_defense": 100, "speed": 100}`
	// pokestatus := `{"hp": "100", "attack": "100", "defense": "100", "special_attack": "100", "special_defense": "100", "speed": "100"}`

	var status Status

	err := json.Unmarshal([]byte(pokestatus), &status)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pokestatus)

	// m := `{"hp": 100, "speed": 100, "attack": 100, "defense": 100, "special_attack": 100, "special_defense": 100}`
	// fmt.Println(`{"hp": 100, "speed": 100, "attack": 100, "defense": 100, "special_attack": 100, "special_defense": 100}`)
	// b, _ := json.Marshal(m)
	// json.NewEncoder(w).Encode(b)
}

func RouterPokemon() {
	http.HandleFunc("/", homePage)

	// Pokemon
	http.HandleFunc("/api/pokemon/show/", middleware.CORS(getPokemonHndler))
	http.HandleFunc("/api/pokemon/list", middleware.CORS(getPokemonsHndler))
	http.HandleFunc("/api/jsondemo", middleware.CORS(getJsonDemoHndler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
