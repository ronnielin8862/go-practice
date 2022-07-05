package main

import "fmt"

type (
	user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	man struct {
		Name string `json:"name"`
	}
	woman struct {
		Name string `json:"name"`
	}
)

func main() {
	man1 := man{Name: "A"}
	woman1 := woman{Name: "B"}
	printNameGeneric(man1)
	printNameGeneric(woman1)

	printNameGenerics([]man{man1})
}

func printNameGenerics[T any](structs []T) {
	fmt.Println("len : ", len(structs))
}

func printName(human string) {
	fmt.Println(human)
}

func printNameGeneric[T any](structs T) {
	fmt.Println(structs)
}
