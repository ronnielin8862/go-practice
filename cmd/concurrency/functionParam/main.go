package main

import (
	"fmt"
	"sync"
	"time"
)

// 以下測試，提供做多線程時參考：
// 當變數宣告為全域，會被其他線程覆蓋，所以要用mutex。  變數宣告在function 內，不會被其他的線程覆蓋

// var a string
func testFunction(s string) {

	var a string
	a = s

	for i := 0; i < 3; i++ {
		fmt.Println(i, "times, a = ", a)

		time.Sleep(2 * time.Second)
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(2)

	go func() {
		testFunction("Hello")
		w.Done()
	}()

	go func() {
		testFunction("World")
		w.Done()
	}()

	w.Wait()
}
