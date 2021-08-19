package main

import (
	"fmt"
)

// func mapFunc[T any, M any](a []T, f func(T) M) []M {
//     n := make([]M, len(a), cap(a))
//     for i, e := range a {
//         n[i] = f(e)
//     }
//     return n
// }

func test1(a []int, f func(int) int) []int {
	fmt.Println("len = ", len(a), ",  cap = ", cap(a))
	n := make([]int, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func main() {
	vi := []int{1, 2, 3, 4, 5, 6}
	// vs := mapFunc(vi, func(v int) int {
	// 	return v*v
	// })
	vs := test1(vi, func(v int) int {
		// fmt.Println(v)
		return v * v
	})
	fmt.Println(vs)
}
