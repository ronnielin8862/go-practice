package main

import (
	"fmt"
)

//
//const (
//	a = iota
//	b
//	c
//	d
//	e
//)

type lineupResp struct {
	Id   int    `json:"lineup"`
	Type string `json:"type"`
}

var ch chan struct{}

func main() {

}

func run(i int) {
	ch <- struct{}{}
	defer func() {
		<-ch
	}()
	fmt.Println("run : ", i)
}
