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
	s := "ABCDE"

	fmt.Println(strings.Index(s, "C"))
}

func testReturn() *lineupResp {
	return nil
}
