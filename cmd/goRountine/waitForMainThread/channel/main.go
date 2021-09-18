package main

import (
	"fmt"
	"time"
)

func say(s string, ch chan<- string) {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
	ch <- s
}

//要主執行緒等待gorountine跑完才結束的三種方法
// 3. channel
func main() {
	ch := make(chan string)

	go say("hello", ch)
	go say("world", ch)

	fmt.Println("dd", <-ch)
	fmt.Println("ww", <-ch)
}
