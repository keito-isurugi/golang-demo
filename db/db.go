package db

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	godotnev "github.com/joho/godotenv"
	"golang_demo/models"
)

var db *gorm.DB

// type User struct {
//   gorm.Model
// 	id int `gorm:"primary_key"`
//   Name  string
// }

func InitDB() {
	connectDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
	fmt.Println("migrate success!")
}

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

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME
	// DBに接続
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	// db.Close()
	fmt.Println("ok!")
}