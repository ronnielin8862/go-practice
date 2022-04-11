package main

import (
	"fmt"
	"testing"
)

var v = "0.12345678901234567890123456789"

func TestFloat(t *testing.T) {

	result := testFloat(v)

	fmt.Println("\n" + "轉換前 = " + v + "\n" + "轉換後 = " + result)

	if result != v {
		t.Error("TestFloat 轉換後不相等")
	} else {

		t.Logf("TestFloat 轉換後相等")
	}
}

func TestRat(t *testing.T) {

	result := testNewRat(v)

	t.Logf("\n" + "轉換前 = " + v + "\n" + "轉換後 = " + result)

	if result != v {
		t.Error("TestRat 轉換後不相等")
	} else {

		t.Logf("TestRat 轉換後相等")
	}
}

func TestShopSpringDecimal(t *testing.T) {
	result := testShopSpringDecimal(v)
	t.Logf("\n" + "轉換前 = " + v + "\n" + "轉換後 = " + result)

	if result != v {
		t.Error("TestShopSpringDecimal 轉換後不相等")
	} else {

		t.Logf("TestShopSpringDecimal 轉換後相等")
	}
}

func TestBigFloat(t *testing.T) {
	result := testBigFloat(v)
	t.Logf("\n" + "轉換前 = " + v + "\n" + "轉換後 = " + result)

	if result != v {
		t.Error("TestBigFloat  轉換後不相等")
	} else {

		t.Logf("TestBigFloat 轉換後相等")
	}

}
