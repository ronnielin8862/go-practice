package main

import (
	"fmt"
	"strings"
)

type lineupResp struct {
	Lineup string `json:"lineup"`
	Type   string `json:"type"`
}

func main() {
	s := "AB,CD,EF"

	fmt.Println(strings.Split(s, ",")[1])
}

func testReturn() *lineupResp {
	return nil
}
