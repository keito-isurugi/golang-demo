package router

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	// "os"
	// _ "github.com/go-sql-driver/mysql"
	// gorm "github.com/jinzhu/gorm"
	// godotnev "github.com/joho/godotenv"
)



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

func Router() {
	// r := NewRouter()
	// http.ListenAndServe(":8080", r)
	// fmt.Println("router!!")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users/", UserHandler)
	// http.HandleFunc("/articles", returnAllArticles)
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