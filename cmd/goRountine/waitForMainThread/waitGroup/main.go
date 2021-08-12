package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)
	go say("hello", wg)
	go say("world", wg)
	wg.Wait()
}
