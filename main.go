// package main

// import (
// 	"fmt"
// 	"golang_demo/router"
// 	"golang_demo/auth"
// 	"golang_demo/demo"
// )

// func main() {
// 	auth.Auth()
// 	router.RouterPokemon()
// 	router.Router()
// 	demo.Hoge()
// }

package main

import "fmt"

func zeroval(ival int) {
    ival = 0
    // fmt.Println("zeroval:", ival)

}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    // fmt.Println("initial:", i)

    // zeroval(i)
    // fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)

		fmt.Println(11233)
}