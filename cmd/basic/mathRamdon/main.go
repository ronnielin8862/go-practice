package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Int31n(40)

	fmt.Println("num = ", num)
}
