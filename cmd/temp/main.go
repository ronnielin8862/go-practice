package main

import (
	"fmt"
	"time"
)

type lineupResp struct {
	Id   int    `json:"lineup"`
	Type string `json:"type"`
}

func main() {
	t := time.Now()
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()

	fmt.Println(today)

	y := time.Now().AddDate(0, 0, -1)
	yesterday := time.Date(y.Year(), y.Month(), y.Day(), 0, 0, 0, 0, t.Location()).Unix()
	fmt.Println(yesterday)

	aa := fmt.Sprintf("select match_id , user_id from db_match_favorite where match_id in ( select match_id from db_matches_schedule where match_time >= %v and match_time < %v )", yesterday, today)
	fmt.Println(aa)
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
