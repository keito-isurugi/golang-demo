package greeting

import "fmt"

func Hello(name string) string {
  message := fmt.Sprintf("Hello, %s", name)
  return message
}


// go mod init github.com/keito-isurugi/golang-demo/module/greeting