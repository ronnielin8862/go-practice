package main

import (
	"testing"
)

//test 遇到undefined 錯誤 可參考https: //blog.csdn.net/helen920318/article/details/105017118
func TestRotateArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	//method1(nums, k)
	//method2(nums, k)
	method1(nums, k)
	t.Log("END!!!!!")
}

// 執行效能分析 go test -bench=. -run=none
func BenchmarkRotateArrayMethod3(b *testing.B) {
	b.ResetTimer()
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	for i := 0; i < b.N; i++ {
		method3(nums, k)
	}
}

func BenchmarkRotateArrayMethod1(b *testing.B) {
	b.ResetTimer()
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	for i := 0; i < b.N; i++ {
		method1(nums, k)
	}
}
