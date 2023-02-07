package router

import (
	"encoding/json"
	"fmt"
	"golang_demo/db"
	"log"
	"net/http"
	"strconv"
	"strings"
	"golang_demo/middleware"
	// "os"
	// _ "github.com/go-sql-driver/mysql"
	// gorm "github.com/jinzhu/gorm"
	// godotnev "github.com/joho/godotenv"
)


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to tha HomaPage!")
	fmt.Println("Endpoint Hit: homePage")
}
func dbMigrate(w http.ResponseWriter, r *http.Request) {
	db.Migrate()
}
func dbSeedUser(w http.ResponseWriter, r *http.Request) {
	db.SeedUser()
}
func dbSeedTodo(w http.ResponseWriter, r *http.Request) {
	db.SeedTodo()
}

func getTodoHndler(w http.ResponseWriter, r *http.Request) {
	todoId := strings.TrimPrefix(r.URL.Path, "/api/todo/show/")
	id, _ := strconv.Atoi(todoId)
	json.NewEncoder(w).Encode(db.GetTodo(id))
}
func getUsersHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetUsers())
}
func getTodosHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetTodos())
}
type Todo struct {
	// jsonで型定義
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}
func registerTodoHndler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 変数を定義
		var todo Todo
		// デコードして変数へ値を格納する
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			log.Fatal(err)
		}
		db.RegisterTodo(todo.Title, todo.Content)
	}
}
func editTodoHndler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			log.Fatal(err)
		}
		db.EditTodo(todo.Id, todo.Title, todo.Content)
	}
}

type TodoId struct {
	Id int `json:"id"`
}
func deleteTodoHndler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var id TodoId
		if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
			log.Fatal(err)
		}
		db.DeleteTodo(id.Id)
	}
}
func todoToUserHndler(w http.ResponseWriter, r *http.Request) {
	db.TodoToUser()
}

type Login struct {
	// jsonで型定義
	Email string `json:"email"`
	Password string `json:"password"`
}
func loginHndler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var login Login
		if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
			log.Fatal(err)
		}
		fmt.Println(db.Login(login.Email, login.Password))
	}
}



// ミドルウェア作成
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is midlleware test!!")
}
func middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("[START] middleware1")
			next.ServeHTTP(w, r)
			fmt.Println("[END] middleware1")
	}
}

func Router() {
	http.HandleFunc("/", homePage)

	// User
	http.HandleFunc("/api/users/", getUsersHndler)

	// Todo
	http.HandleFunc("/api/todo/list", middleware.CORS(getTodosHndler))
	http.HandleFunc("/api/todo/show/", middleware.CORS(getTodoHndler))
	http.HandleFunc("/api/todo/register", middleware.CORS(registerTodoHndler))
	http.HandleFunc("/api/todo/edit", middleware.CORS(editTodoHndler))
	http.HandleFunc("/api/todo/delete", middleware.CORS(deleteTodoHndler))
	http.HandleFunc("/api/todo_to_user", middleware.CORS(todoToUserHndler))

	// 認証関係
	// http.HandleFunc("/api/login", middleware.CORS(loginHndler))
	
	// マイグレーション, シード
	http.HandleFunc("/db/migrate", dbMigrate)
	http.HandleFunc("/db/seed/user", dbSeedUser)
	http.HandleFunc("/db/seed/todo", dbSeedTodo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	i := strings.Index(path, "/")
	userId, path := path[:i], path[i:]
	
	fmt.Println("users userId:", userId)
	// fmt.Println("users path:", path)
	PostHndler(w, r, path, userId)
}

func PostHndler(w http.ResponseWriter, r *http.Request, path string, userId string) {
	path = strings.TrimPrefix(path, "/posts/")
	i := strings.Index(path, "/")

	var postId string
	if(i <= 0) {
		postId = path
	} else {
		postId, path = path[:i], path[i:]
	}

	// CommentHandler(w, r, path, userId, postId)
	// fmt.Println("posts path:", path)
	fmt.Println("posts postId:", postId)
}