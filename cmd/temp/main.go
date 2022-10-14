package main

import (
	"fmt"
)

type lineupResp struct {
	Id   int    `json:"lineup"`
	Type string `json:"type"`
}

func main() {
	fmt.Println("2222 = ", 50/100)
}

func testReturn(s string) string {
	b := []byte(s)
	fmt.Println("len :", len(b))
	j := len(b) - 1
	for i := 0; i < len(b)/2; i++ {
		b[i], b[j] = b[j], b[i]
		j--
	}
	return string(b)
}
