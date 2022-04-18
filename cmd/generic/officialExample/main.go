package main

import (
	"fmt"
)

func main() {
	//generic map
	ints := map[string]int64{"first": 10, "second": 20}
	floats := map[string]float64{"first": 0.1, "second": 0.2}
	fmt.Printf("numFromMap ： ints = %v , floats = %v \n", numFromMap(ints), numFromMap(floats))

	//generic slice
	intArray := []int64{1, 2, 3, 4}
	floatArray := []float64{0.1, 0.2, 0.3}
	fmt.Printf("numFromSlice ： ints = %v , floats = %v \n", numFromSlice(intArray), numFromSlice(floatArray))

	//傳入相同類型，返回不同類型
	fmt.Println("stings a = ", returnGeneric("a"))
	fmt.Println("stings b = ", returnGeneric("b"))
}

func numFromMap[k comparable, V int64 | float64](m map[k]V) (s V) {
	for _, v := range m {
		s += v
	}
	return s
}

func numFromSlice[V int64 | float64](m []V) (result V) {
	for _, v := range m {
		result += v
	}
	return result
}

func returnGeneric(arg string) any {
	if arg == "a" {
		a := []string{"a", "b"}
		return a
	}
	return 1234567
}

//comparable 必須是與相同的參數類型(V) 才能match type
//func returnGeneric2[V comparable](ggg V, arg string) V {
//	if arg == "a" {
//		a := []string{"a", "b"}
//		//返回其他comparable 類型就會報錯
//		return a
//	}
//
//	if ggg == ggg {
//
//	}
//	return ggg
//}

func find[V comparable](what V, s []V) int {
	for i, v := range s {
		if v == what {
			return i
		}
	}
	return -1
}

type MyMap[T comparable, V int64 | float64] map[T]V
