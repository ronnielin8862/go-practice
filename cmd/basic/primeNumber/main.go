package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	loopNum := 200000

	forLoopFunc(loopNum)

	fmt.Println("=====================================")

	goRoutineFunc(loopNum)
}

func goRoutineFunc(loopNum int) {
	defer timeCost()()
	c := make(chan int)
	primeChan := make(chan int, loopNum)

	for i := 0; i < 9; i++ {
		go startGoRoutine(c, primeChan)
	}

	for i := 2; i < loopNum; i++ {
		c <- i
	}
	close(c)
	fmt.Println("goRoutine func result : num of prime ", len(primeChan))
}

func startGoRoutine(c chan int, primeChan chan int) {
	for {
		if isPrime(<-c) {
			primeChan <- 1
		}
	}
}

func forLoopFunc(loopNum int) {
	defer timeCost()()
	var list []int
	for i := 0; i < loopNum; i++ {
		if isPrime(i) {
			list = append(list, i)
		}
	}
	fmt.Println("for loop func result : num of prime ", len(list))
}

func isPrime(num int) bool {
	for j := 2; j < num; j++ {
		if num%j == 0 {
			return false
		}
	}
	return true
}

func timeCost() func() {
	startTime := time.Now()
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return func() {
		t := time.Since(startTime)
		fmt.Printf("%s time cost : %v /ms \n", f.Name(), t.Milliseconds())
	}
}
