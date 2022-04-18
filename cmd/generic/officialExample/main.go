package main

import (
	"fmt"
)

func main() {
	ints := map[string]int64{"first": 10, "second": 20}
	floats := map[string]float64{"first": 0.1, "second": 0.2}
	fmt.Printf("genericNumFromMap ： ints = %v , floats = %v \n", genericNumFromMap(ints), genericNumFromMap(floats))

	intArray := []int64{1, 2, 3, 4}
	floatArray := []float64{0.1, 0.2, 0.3}
	fmt.Printf("genericNumFromSlice ： ints = %v , floats = %v \n", genericNumFromSlice(intArray), genericNumFromSlice(floatArray))

	//傳入相同類型，返回不同類型
	fmt.Println("stings a = ", returnGeneric("a"))
	fmt.Println("stings b = ", returnGeneric("b"))

}

func genericNumFromMap[k comparable, V int64 | float64](m map[k]V) (s V) {
	for _, v := range m {
		s += v
	}
	return s
}

func genericNumFromSlice[V int64 | float64](m []V) (result V) {
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

//
//func genericString[V int](m V) []any {
//	var s []string
//
//	if reflect.TypeOf(m)
//	aaaa := len(m)
//	for _, v := range len(m) {
//		s = append(s, string(v))
//	}
//	return s
//}
//
//func test(v int) {
//	aa := len()
//}

//var gg g[int64]

func testArg(arg MyMap[int, int64]) {

}

type MyMap[T comparable, V int64 | float64] map[T]V
