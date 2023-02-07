package db

import (
	// "encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	godotnev "github.com/joho/godotenv"
	// "golang_demo/models"
)

var db *gorm.DB

// db接続
func ConnectDB() {
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
