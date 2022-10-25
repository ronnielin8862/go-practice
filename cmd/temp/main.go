package main

import (
	"fmt"
)

type lineupResp struct {
	Id   int    `json:"lineup"`
	Type string `json:"type"`
}

var ch chan struct{}

func main() {
	m := make(map[int]lineupResp)
	m[1] = lineupResp{Id: 1, Type: "1"}
	m[2] = lineupResp{Id: 2, Type: "2"}

	fmt.Println("m 3: ", m[3])
}

func run(i int) {
	ch <- struct{}{}
	defer func() {
		<-ch
	}()
	fmt.Println("run : ", i)
}
