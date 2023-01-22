package models

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	godotnev "github.com/joho/godotenv"
	"golang_demo/typefile"
)

var db *gorm.DB

// type User struct {
//   gorm.Model
// 	id int `gorm:"primary_key"`
//   Name  string
// }

func InitDB() {
	connectDB()
	db.AutoMigrate(&typefile.User{})
	db.AutoMigrate(&typefile.Todo{})
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