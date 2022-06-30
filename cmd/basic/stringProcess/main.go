package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var users1 []user
	var users2 []user
	loop := 100000
	for i := 0; i < loop; i++ {
		users1 = append(users1, user{id: i, name: "name" + strconv.Itoa(i)})
	}
	copy(users2, users1)

	t1 := time.Now().UnixMilli()

	for _, v := range users1 {
		for _, j := range users2 {
			if v.id == j.id {
				continue
			}
		}
	}

	t2 := time.Now().UnixMilli()

	fmt.Println(" 2 for loops : ", t2-t1)

	t3 := time.Now().UnixMilli()

	s := make(map[int]int)
	for k, v := range users1 {
		s[v.id] = k
	}
	for _, v := range users2 {
		if _, ok := s[v.id]; !ok {
			fmt.Println(" !ok ")
		}
		user := users1[s[v.id]]
		if user.id == v.id {
			continue
		}
	}
	t4 := time.Now().UnixMilli()

	fmt.Println(" map : ", t4-t3)
}

type user struct {
	id   int    `json:"id"`
	name string `json:"name"`
}
