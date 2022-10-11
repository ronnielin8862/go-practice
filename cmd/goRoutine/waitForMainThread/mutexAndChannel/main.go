package main

import (
	"fmt"
	"sync"
	"time"
)

func addByShareMemory(n int, ch chan string) []int {
	var ints []int
	var mux sync.Mutex

	for i := 0; i < n; i++ {
		go func(i int) {
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()
		}(i)
	}
	time.Sleep(time.Second)

	ch <- "go"
	return ints
}

/*
* TODO:
* 2021/8/12: 使用channel做阻塞，遇到死結問題待處理
* 2021/8/13: 為channel增加指定數量後即可解決，Y????
 */
func main() {

	ch := make(chan string, 1)
	foo := addByShareMemory(10, ch)
	fmt.Println(<-ch)
	fmt.Println(len(foo))
	fmt.Println(foo)

}
