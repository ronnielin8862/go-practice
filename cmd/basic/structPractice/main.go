package main

import "fmt"

type APlusB func(int, int) int

type Animal struct {
	name     string
	behavior string
	aPlusB   APlusB
}

func (a *APlusB) Do(b int) int {
	fmt.Println(b)
	return 1 + 1
}

func (a *Animal) Do1(b int) int {
	fmt.Println(b)
	return 2 + 2
}

func main() {
	//animal := Animal{
	//	name:     "cat",
	//	behavior: "meow",
	//	aPlusB:   func(a int, b int) int { return a + b },
	//}
	//
	//animal.Do1(9 + 9)
	//
	//b := animal.aPlusB(3, 3)
	//fmt.Println(b)

	//fmt.Println(animal.aPlusB.Do(4))

	// 此例可得知，golang struct 比較的是內函數所有的值
	compareStruct()
}

func compareStruct() {
	aStruct := A{
		name: "a",
		Z:    Z{num: 1},
	}

	bStruct := A{
		name: "b",
		Z:    Z{num: 2},
	}

	aStruct2 := A{
		name: "a",
		Z:    Z{num: 1},
	}

	aStruct3 := A{
		name: "a",
		Z:    Z{num: 2},
	}
	aStruct4 := A{
		name: "a",
		Z:    Z{num: 1},
	}

	fmt.Println(aStruct == bStruct)
	fmt.Println(aStruct == aStruct2)
	fmt.Println(aStruct == aStruct3)
	fmt.Println(aStruct == aStruct4)
}

type (
	Z struct {
		num int
	}

	A struct {
		name string
		Z
	}
)
