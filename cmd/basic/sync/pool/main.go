package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type man struct {
	name   string
	age    int
	Remark [1024]byte
}

var manJson, _ = json.Marshal(man{name: "abc", age: 1})

func main() {
	pool()
}

func pool() {
	var p sync.Pool
	var ms []man
	p.New = func() interface{} {
		return new(man)
	}

	m := p.Get().(*man)
	m.name = "BBB"

	for i := 0; i <= 2; i++ {
		m := p.Get().(*man)
		m.age = i
		ms = append(ms, *m)
	}

	fmt.Println("r =", ms)
}
