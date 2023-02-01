package demo

import (
	"fmt"
	// "math/rand"
	// "sync"
	// "time"
)

func GetHello(s string) string {
  return "こんにちは、" + s + "！"
}

func Hello() {
	s := GetHello("山澤さん")
  fmt.Println(s)
}