package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

//要主執行緒等待gorountine跑完才結束的三種方法

// 1. 指定等待時間
func main() {
	go say("hello")
	say("world")

	time.Sleep(time.Second * 2)
}
