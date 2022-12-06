package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d, e = iota, iota * 10
	f, g
	h, i
)

func main() {
	fmt.Println(a, b, c)

	fmt.Println(d, e, f, g, h, i)
}
