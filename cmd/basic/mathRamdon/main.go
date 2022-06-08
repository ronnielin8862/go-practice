package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand_Float64()
}

func rand_Float64() {
	fmt.Println(rand.Float64())
}

func rand_Int31n() {

	num := rand.Int31n(40)

	fmt.Println("num = ", num)
}
