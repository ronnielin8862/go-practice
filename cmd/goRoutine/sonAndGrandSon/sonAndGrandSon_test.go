package sonAndGrandSon

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 子协程和孙协程： 子協程結束後，孫協程還沒結束時，孫協程會繼續執行。 與主協程結束後子協程會被關閉的情況不同
func Test1(t *testing.T) {
	r := sonAndGranSon()
	time.Sleep(2 * time.Second)
	fmt.Println("r : ", r)
}

func sonAndGranSon() []int {
	var r []int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {

		go func() {
			time.Sleep(1 * time.Second)
			r = append(r, 2)
			wg.Done()
		}()

		r = append(r, 1)
		wg.Done()
	}()

	wg.Wait()
	r = append(r, 3)
	return r
}
