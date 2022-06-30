package main

import (
	"fmt"
)

func main() {
	u1 := user{id: 1}
	u2 := user{id: 2}
	us := []user{u1, u2}
	fmt.Println(us)
}

type user struct {
	id int
}
