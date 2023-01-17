package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
	godotnev "github.com/joho/godotenv"
)

type Player struct {
	// gorm.Model
	Id int
	Name string
}

// var player Player
// var players []Player



// type Article struct {
// 	Title string `json: "Title"`
// 	Desc string `json: "desc"`
// 	Content string `json: "content"`
// }

// type Articles []Article

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

// func fetchPlayers(w http.ResponseWriter, r *http.Request) {
// 	db := GetDbConn()

// 	db.Find(&players)
// 	fmt.Println(players)
// 	profJson, _ := json.Marshal(players)
// 	fmt.Fprintf(w, string(profJson))
// }

// func GetDbConn() *gorm.Db {
// 	db, err := gorm.Open(GetDBConfig())
// 	if err != nil {panic(err)}

// 	db.LogModel(true)
// 	return db
// }

// func GetDBConfig() (string, string) {
// 	DBMS := "mysql"
// 	USER := "test"
// 	PASS := "test"
// 	PROTOCOL := ""
// 	DBNAME := "golang_test"
// 	OPTION := "charset=utf8&parseTime=True&loc=Local"
// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

// 	return DBMS, CONNECT
// }

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



func main() {
	// ENV読み取り
	err := godotnev.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}

	// 接続情報を設定
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	// (localhost:3306ではなく) (コンテナ名:3306)
	HOST := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME

	// DBに接続
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	// var player Player
	var players []Player

	db.Find(&players)
	defer db.Close()

	fmt.Print(players)
	handleReauests()
}