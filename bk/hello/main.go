package main

import (
  "fmt"

	"github.com/keito-isurugi/golang-demo/module/greeting"
)

func main() {
  message := greeting.Hello("John")
  fmt.Println(message)
}


// go mod init github.com/keito-isurugi/golang-demo/module/hello
