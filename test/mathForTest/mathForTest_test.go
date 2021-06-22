package main

import (
	"fmt"
	"testing"

	"github.com/ronnielin8862/go-api/pkg/mathForTest"
)

func TestMathPlus(t *testing.T) {
	fmt.Println("開始測試加法")

	var a, b = 2, 3

	result := mathForTest.Plus(a, b)

	if result != a+b {
		t.Error("加法錯誤")
	}
}
