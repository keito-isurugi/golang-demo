package db

import (
	// "encoding/json"
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	godotnev "github.com/joho/godotenv"
	"golang_demo/models"
)

var db *gorm.DB

// db接続
func connectDB() {
	// ENV読み取り
	err := godotnev.Load(".env")
	if err != nil {
		fmt.Printf("読み込みに失敗しました: %v", err)
	}

	// 接続情報を設定
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	// (localhost:3306ではなく)
	HOST := "tcp(golang_demo_db:3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"
	// DBに接続
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	// db.Close()
	fmt.Println("ok!")
}


// マイグレーション
func Migrate() {
	connectDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
	fmt.Println("migrate success!")
}


// データ登録
func SeedUser() {
	connectDB()
	var user models.User
	for i := 0; i < 10; i++ {
		user = models.User{Name: fmt.Sprintf("user_%v", i)}
		db.Create(&user)
	}
	defer db.Close()
	fmt.Println("seeder success!")
}

func SeedTodo() {
	connectDB()
	var todo models.Todo
	for i := 0; i < 10; i++ {
		todo = models.Todo{Title: fmt.Sprintf("タイトル_%v", i), Content: fmt.Sprintf("内容_%v", i)}
		db.Create(&todo)
	}
	defer db.Close()
	fmt.Println("seeder success!")
}


// データ取得
func GetUsers() []models.User{
	var users []models.User
	connectDB()
	db.Find(&users)
	defer db.Close()
	return users
}
func GetTodos() []models.Todo{
	var todos []models.Todo
	connectDB()
	db.Find(&todos)
	defer db.Close()
	return todos
}

type Todo struct {
	// jsonで型定義
	Title string `json:"title"`
	Content string `json:"content"`
}

// データ登録
func RegisterTodo(title string, content string) {
	connectDB()
	todo := Todo{Title: title, Content: content}
	db.Create(&todo)
	defer db.Close()
	fmt.Println("Todoを登録しました")
}

// データを削除
func DeleteTodo(id int) {
	var todo models.Todo
	connectDB()
	db.Delete(&todo, id)
}
