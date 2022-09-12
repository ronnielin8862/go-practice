package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	//maxLen, start := 0, 0
	table := [128]int{}
	for i, _ := range table {
		table[i] = -1
	}
	fmt.Println(table)
}
