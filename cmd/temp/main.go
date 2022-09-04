package main

import (
	"bytes"
	"fmt"
)

type lineupResp struct {
	Lineup string `json:"lineup"`
	Type   string `json:"type"`
}

func main() {
	var s []string
	s = append(s, "")

	fmt.Println("aaaa", s[0])
}

func lineupToJason(lineup string) string {
	var buffer bytes.Buffer

	buffer.WriteString("{\"linup1\":\"")
	buffer.WriteString(lineup)
	buffer.WriteString("\"}")
	return buffer.String()
}
