package addTwoNumbers

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	var list1 []int
	var list2 []int
	var aList []int
	list1 = append(list1, 2, 3, 4)
	list2 = append(list2, 5, 6, 4)
	aList = append(aList, 7, 9, 8)

	a := addTwoNumbers(list1, list2)
	t.Log(a)

	if reflect.DeepEqual(a, aList) {
		t.Log("1 success")
	} else {
		t.Error("1 fail")
	}
}

func Test2(t *testing.T) {
	var list1 []int
	var list2 []int
	var aList []int
	list1 = append(list1, 0, 3, 9)
	list2 = append(list2, 0, 6, 9)
	aList = append(aList, 0, 9, 8)

	a := addTwoNumbers(list1, list2)
	t.Log(a)

	if reflect.DeepEqual(a, aList) {
		t.Log("2 success")
	} else {
		t.Error("2 fail")
	}
}
