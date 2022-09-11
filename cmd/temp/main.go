package main

import (
	"fmt"
)

type lineupResp struct {
	Lineup string `json:"lineup"`
	Type   string `json:"type"`
}

func main() {
	fmt.Println("1 :", &lineupResp{})
	b := &lineupResp{}
	fmt.Println("2 :", b)

	c := testReturn()
	fmt.Println("3 :", c)

	d := testReturn()
	fmt.Println("4 :", d.Lineup)
}

func testReturn() *lineupResp {
	return nil
}
