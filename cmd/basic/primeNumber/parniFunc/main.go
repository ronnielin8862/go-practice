package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 200000; i++ {
		intChan <- i
	}
	close(intChan)
}
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("协程取不到数据 执行完毕")
	exitChan <- true
}
func main() {
	intChan := make(chan int, 20000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 10)
	go putNum(intChan)
	for i := 0; i < 10; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 10; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
	fmt.Println("方法执行完毕")
}
