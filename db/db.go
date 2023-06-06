package db

import (
	// "encoding/json"
	"fmt"
	"golang_demo/models"

	_ "github.com/go-sql-driver/mysql"
)

// マイグレーション
func MigrateTodo() {
	ConnectDB()
	db.AutoMigrate(&models.Todo{})
	fmt.Println("migrate todo success!")
}

// データ登録
func SeedTodo() {
	ConnectDB()
	var todo models.Todo
	for i := 1; i <= 10; i++ {
		todo = models.Todo{UserID: 1, Title: fmt.Sprintf("タイトル_%v", i), Content: fmt.Sprintf("内容_%v", i)}
		db.Create(&todo)
	}
	defer db.Close()
	fmt.Println("seeder success!")
}

// データ取得
func GetTodos() []models.Todo {
	var todos []models.Todo
	ConnectDB()
	db.Preload("User").Find(&todos)
	defer db.Close()
	return todos
}
func GetTodo(id int) models.Todo {
	var todo models.Todo
	ConnectDB()
	db.First(&todo, id)
	defer db.Close()
	return todo
}

type Todo struct {
	// jsonで型定義
	Title   string `json:"title"`
	Content string `json:"content"`
}

// データ登録
func RegisterTodo(title string, content string) {
	ConnectDB()
	todo := Todo{Title: title, Content: content}
	db.Create(&todo)
	defer db.Close()
	fmt.Println("Todoを登録しました")
}

// データ更新
func EditTodo(id int, title string, content string) {
	var todo models.Todo
	ConnectDB()
	db.First(&todo, id)
	db.Model(&todo).Updates(models.Todo{Title: title, Content: content})
	defer db.Close()
	fmt.Println("Todoを更新しました")
}

// データを削除
func DeleteTodo(id int) {
	var todo models.Todo
	ConnectDB()
	db.Delete(&todo, id)
}

// リレーションお試し
func TodoToUser() {
	// var todo models.Todo
	var users []models.User
	ConnectDB()
	db.Preload("Todos").Find(&users)
	// db.Find(&users)
	fmt.Println(users)
}
