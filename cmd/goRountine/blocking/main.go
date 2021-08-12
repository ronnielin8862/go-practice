package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println("A = ", v)
		sum += v
	}
	c <- sum

	wg.Done()
}

func sum2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println("B = ", v)

		sum += v
	}
	c <- sum

	wg.Done()
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(s[:4])
	wg.Add(2)

	c := make(chan int, 2)

	go sum1(s, c)
	go sum2(s, c)

	wg.Wait()

}
