package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var t int = 50000
	// 簡略版本
	fmt.Printf("n:%v \n\n", sum(t))

	// 優美版本
	fmt.Println("n2: ", sum2(t))
}

func sum(n int) int {
	// defer 會先進入執行，待離開function時返回
	defer fmt.Println("defer return: ", timeCost())
	a := time.Now()

	var s int = 0
	for i := 0; i < n; i++ {
		s += i
	}
	sc := time.Since(a)
	println("sum time cost ", sc.Nanoseconds())
	return s
}

func timeCost() int {
	fmt.Println("start")
	return 999
}

func sum2(n int) int {
	defer timeCost2()() // 不加上第二個"()", 只回返回方法卻不執行。 如果print timeCost2()，會得到function address。
	var s int = 0
	for i := 0; i < n; i++ {
		s += i
	}
	return s
}

func timeCost2() func() {
	startTime := time.Now()
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return func() {
		t := time.Since(startTime)
		fmt.Println(f.Name(), " time cost : ", t.Nanoseconds())
	}
}
