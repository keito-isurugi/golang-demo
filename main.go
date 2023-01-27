package main

import (
	"golang_demo/router"
	"golang_demo/auth"
)

func main() {
	auth.Auth()
	router.Router()
}