package main

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang_demo/router"
	// "golang_demo/db"
	// "golang_demo/models"
)

func main() {
	router.Router()
	// models.InitDB()
}


// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func helloHander(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Hello Normal isurugi</h1>")
// }

// func main() {
// 	http.HandleFunc("/", helloHander)
// 	http.ListenAndServe(":8080", nil)
// }
