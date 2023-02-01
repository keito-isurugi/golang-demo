package demo

import (
	"fmt"
	"math/rand"
	"sync"

	// "sync"
	"time"
)

func GetLuckyNum(c chan <- int) {
	fmt.Println("...")
	
	rand.Seed((time.Now().Unix()))
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	// fmt.Printf("Today's your lucky number is %d!\n", num)

	c <- num
}

func GoRoutineMain() {
	// for i := 0; i < 3; i++ {
  //   go func(i int) {
  //       fmt.Println(i)
  //   }(i)
	// }

	// fmt.Println("what is today's lucky number?")

	// c := make(chan int)
	// go GetLuckyNum(c)

	// num := <-c
	// fmt.Printf("Today's your lucky number is %d!\n", num)
	// close(c)

	// time.Sleep(time.Second * 5)


	// 待ち合わせあり
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	GetLuckyNum()
	// }()

	// wg.Wait()

	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	// c := make(chan int)
	// // srcの要素毎にある何か処理をして、結果をdstにいれる
	// for _, s := range src {
	// 	go func(s int, c chan int) {
	// 		// 何か(重い)処理をする
	// 		result := s * 2
	// 		c <- result
	// 		// 結果をdstにいれる
	// 		// dst = append(dst, result)
	// 	}(s, c)
	// }

	// for _ = range src {
	// 	num := <- c
	// 	dst = append(dst, num)
	// }

	var mu sync.Mutex

	for _, s := range src {
		go func(s int) {
			result := s * 2
			mu.Lock()
			dst = append(dst, result)
			mu.Unlock()
		}(s)
	}
	time.Sleep(time.Second)
	fmt.Println(dst)
	// close(c)
}