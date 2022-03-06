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
	animal := Animal{
		name:     "cat",
		behavior: "meow",
		aPlusB:   func(a int, b int) int { return a + b },
	}

	animal.Do1(9 + 9)

	b := animal.aPlusB(3, 3)
	fmt.Println(b)

	fmt.Println(animal.aPlusB.Do(4))
}
