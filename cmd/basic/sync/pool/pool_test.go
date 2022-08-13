package main

import (
	"encoding/json"
	"sync"
	"testing"
)

func BenchmarkSyncPool(b *testing.B) {
	var p sync.Pool
	p.New = func() interface{} {
		return new(man)
	}

	for i := 0; i <= 10000; i++ {
		m := p.Get().(*man)
		json.Unmarshal(manJson, m)
		p.Put(m)
	}
}

func BenchmarkSyncPool2(b *testing.B) {
	for i := 0; i <= 10000; i++ {
		//var m man
		m := &man{}
		json.Unmarshal(manJson, m)
	}
}

//func BenchmarkTest(b *testing.B) {
//	num := 10
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		fmt.Sprintf("%d", num)
//	}
//}
