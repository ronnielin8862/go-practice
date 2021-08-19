package main

import "fmt"

func test1() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[2:6]

	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}

func test2() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[5:8]

	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}

func test3() {
	a := make([]int, 10, 20)

	fmt.Println(a)
}

func main() {
	// test1()
	// test2()
	test3()
}
