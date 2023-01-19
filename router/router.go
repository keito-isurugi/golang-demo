package router

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"

	// _ "github.com/go-sql-driver/mysql"
	// gorm "github.com/jinzhu/gorm"
	// godotnev "github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to tha HomaPage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleReauests() {
	http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", returnAllArticles)
	// http.HandleFunc("/players", fetchPlayers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	articles := Articles{}
// 	for i := 0; i < 10; i++ {
// 			title := "Hello_%d"
// 			articles = append(
// 					articles,
// 					Article{Title: fmt.Sprintf(title, i), Desc: "Article Description", Content: "Article Content"})
// 	}
// 	fmt.Println("Endpoint Hit: returnAllArticles")
// 	json.NewEncoder(w).Encode(articles)
// }

func Router() {
	fmt.Println("router!!")
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", returnAllArticles)
	// http.HandleFunc("/players", fetchPlayers)
	// log.Fatal(http.ListenAndServe(":8081", nil))
}