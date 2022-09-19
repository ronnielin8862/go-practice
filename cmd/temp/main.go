package main

import "fmt"

type lineupResp struct {
	Lineup string `json:"lineup"`
	Type   string `json:"type"`
}

func main() {

	s := "AB,CD,EF"
	fmt.Println(testReturn(s))
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
