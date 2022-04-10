package main

import (
	"fmt"
	"testing"
)

var v = "0.12345678901234567890123456789"

func TestFloat(t *testing.T) {

	result := testFloat(v)

	fmt.Println("轉換前 = " + v + "\n" + "轉換後 = " + result)

	if result != v {
		t.Error("TestFloat轉換後不相等")
	}
}

func TestRat(t *testing.T) {

	result := testNewRat(v)

	t.Logf("轉換前 = " + v + "\n" + "轉換後 = " + result)

	if result != v {
		t.Error("TestRat轉換後不相等")
	}
}

func TestShopSpringDecimal(t *testing.T) {
	result := testShopSpringDecimal(v)
	t.Logf("轉換前 = " + v + "\n" + "轉換後 = " + result)

	if result != v {
		t.Error("TestShopSpringDecimal轉換後不相等")
	}

	t.Logf("TestShopSpringDecimal 轉換後相等")
}
