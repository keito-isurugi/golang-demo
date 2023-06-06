package db

import (
	// "encoding/json"
	"fmt"
	"golang_demo/models"

	_ "github.com/go-sql-driver/mysql"
)

// マイグレーション
func MigrateUser() {
	ConnectDB()
	db.AutoMigrate(&models.User{})
	fmt.Println("migrate user success!!")
}

// データ登録
func SeedUser() {
	ConnectDB()
	var user models.User
	for i := 1; i <= 10; i++ {
		user = models.User{
			Name:     fmt.Sprintf("user_%v", i),
			Email:    fmt.Sprintf("test%v@email.com", i),
			Password: fmt.Sprintf("test%v", i),
		}
		db.Create(&user)
	}
	defer db.Close()
	fmt.Println("seeder success!")
}

// データ取得
func GetUsers() []models.User {
	var users []models.User
	ConnectDB()
	db.Find(&users)
	defer db.Close()
	return users
}
func GetUser(id int) models.User {
	var user models.User
	ConnectDB()
	db.First(&user, id)
	defer db.Close()
	return user
}

func GetLoginUser(email string, password string) models.User {
	var user models.User
	ConnectDB()
	db.Where("email = ?", email).Where("password = ?", password).First(&user)
	defer db.Close()
	return user
}

// ログイン
func Login(email string, password string) models.User {
	var user models.User
	ConnectDB()
	db.Where("email = ?", email).Where("password = ?", password).First(&user)
	return user
}
