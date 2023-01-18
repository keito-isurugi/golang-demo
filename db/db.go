package db

import (
	// "encoding/json"
	"fmt"
	// "log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
	godotnev "github.com/joho/godotenv"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Users User

func Db() {
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
	HOST := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME

	// DBに接続
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	// データ取得
	var user User
	// db.First(&user)
	// fmt.Println(user)

	// var users []Users
	// db.Find(&users)
	// defer db.Close()
	// fmt.Print(users)

	// データ登録
	// for i := 4; i <= 10; i++ {
	// 	user := User{ID: i, Name: fmt.Sprint("渡辺_", i), Age: 10 + i}
	// 	db.Create(&user)
	// }
	// user := User{ID: 3, Name: "suzuki", Age: 18}
	// result := db.Create(&user)

	// fmt.Println(result.Error)
	// fmt.Println(user.ID)
	// fmt.Println(result.RowsAffected)

	// レコードの更新
		// user.ID = 1
		// user.Name = "jinzhu 2"
		// user.Age = 100
		// db.Save(&user)
		
		// db.Model(&User{}).Where("id = ?", 1).Update("name", "hello")
		
		// データ削除
		db.Where("name = ?", "渡辺_4").Delete(&user)
}
