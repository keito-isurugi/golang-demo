package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"golang_demo/db"
	// "golang_demo/middleware"
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
func getUsersHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetUsers())
}
func getTodosHndler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.GetTodos())
}
func registerTodoHndler(w http.ResponseWriter, r *http.Request) {
	db.RegisterTodo()
}


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


// セッションの確認
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// クロスオリジン用にセット
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")
    w.Header().Set("Content-Type", "application/json")

		// preflight用に200でいったん返す
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func Router() {
	// r := NewRouter()
	// http.ListenAndServe(":8080", r)
	// fmt.Println("router!!")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/users/", getUsersHndler)
	http.HandleFunc("/api/todo/list", CORS(getTodosHndler))
	// http.HandleFunc("/api/todo/list", getTodosHndler)
	http.HandleFunc("/api/todo/register", registerTodoHndler)
	
	http.HandleFunc("/db/migrate", dbMigrate)
	http.HandleFunc("/db/seed/user", dbSeedUser)
	http.HandleFunc("/db/seed/todo", dbSeedTodo)
	// http.HandleFunc("/players", fetchPlayers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// type Router struct{}

// func NewRouter() http.Handler {
// 	r := &Router{}
// 	return r
// }

// func (router *Router) ServerHTTP(w http.ResponseWriter, r *http.Request) {
// endpoint, param := SeparatePath(...)

// 	switch endpoint {
// 	case "/users":
// 		switch r.Method {
// 		case http.MethodGet:
// 			UserHandler(w, r, param)
// 		default:
// 			MethodNotAllowedHandler(w, r)
// 		}
// 	default:
// 		NotFoundHandler(w, r)
// 	}
// }

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