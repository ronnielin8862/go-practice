package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type lineupResp struct {
	Lineup string `json:"lineup"`
	Type   string `json:"type"`
}

func main() {
	var a []string = []string{"123456789x"}
	b := a[0][:len(a[0])/3]
	c := a[0][len(a[0])/3 : len(a[0])/3*2]
	d := a[0][len(a[0])/3*2:]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(lineupToJason(b))
	fmt.Println(lineupToJason(c))
	fmt.Println(lineupToJason(d))

	var resp lineupResp
	resp.Lineup = b
	resp.Type = "lineup1"

	rrr, _ := json.Marshal(resp)
	fmt.Println(string(rrr))
}

func lineupToJason(lineup string) string {
	var buffer bytes.Buffer

	buffer.WriteString("{\"linup1\":\"")
	buffer.WriteString(lineup)
	buffer.WriteString("\"}")
	return buffer.String()
}
