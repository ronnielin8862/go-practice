package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mut sync.Mutex
	v := 0

	for i := 0; i < 1000; i++ {
		go func() {
			mut.Lock()
			v++
			mut.Unlock()
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println(v)
}
