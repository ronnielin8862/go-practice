package main

import (
	"fmt"
)

type a struct {
	a int
}

//
//const (
//	a = iota
//	b
//	c
//	d
//	e
//)

type b struct {
	b int
}

type lineupResp struct {
	Id   int    `json:"lineup"`
	Type string `json:"type"`
}

type c struct {
	c int
}

var ch chan struct{}

func main() {
	fmt.Println("hello world")
	fmt.Println("hello world")
}

type d struct {
	d int
}

func run(i int) {
	ch <- struct{}{}
	defer func() {
		<-ch
	}()
	fmt.Println("run : ", i)
}
