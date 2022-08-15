package main

import (
	"sync"
	"testing"
)

func BenchmarkSyncPool(b *testing.B) {
	var p sync.Pool
	p.New = func() interface{} {
		return new(man)
	}

	for i := 0; i <= 36300000; i++ {
		m := p.Get().(*man)
		m.name = "AAA"
		m.age = i
		p.Put(m)
	}
}

func BenchmarkSyncPool2(b *testing.B) {
	for i := 0; i <= 36300000; i++ {
		m := new(man)
		//m := &man{}
		//json.Unmarshal(manJson, m)
		m.name = "BBB"
		m.age = i
		//fmt.Println(&m)
	}
}

//func BenchmarkTest(b *testing.B) {
//	num := 10
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		fmt.Sprintf("%d", num)
//	}
//}
