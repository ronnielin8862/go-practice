package firstTest

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestFirst(t *testing.T) {
	var i = []int{25, 2, 3, 57, 38, 41}
	var r []int
	r = test2(i)
	assert.Equal(t, []int{2, 3, 5}, r, "Test1 they should be equal, your answer is : %v", r)

	var i2 = []int{11, 22, 34, 56, 999, 10}
	var r2 []int
	r2 = test2(i2)
	assert.Equal(t, []int{1, 9}, r2, "Test1 they should be equal, your answer is : %v", r2)

}
func test2(a []int) []int {
	m := make(map[string]int)
	var c []int
	maxCount := 0
	for _, v := range a {
		for _, v2 := range strconv.Itoa(v) {
			m[string(v2)] = m[string(v2)] + 1

			if m[string(v2)] > maxCount {
				maxCount = m[string(v2)]
				c = []int{}
				num, _ := strconv.Atoi(string(v2))
				c = append(c, num)
			} else if m[string(v2)] == maxCount {
				num, _ := strconv.Atoi(string(v2))
				c = append(c, num)
			}
		}
	}

	b := quickSort(c)
	return b
}

func test1(a []int) []int {
	m := make(map[string]int)
	for _, v := range a {
		for _, v2 := range strconv.Itoa(v) {
			m[string(v2)] = m[string(v2)] + 1
		}
	}
	a = []int{}
	maxCount := 0
	for k, v := range m {
		if v > maxCount {
			maxCount = v
			a = []int{}
			num, _ := strconv.Atoi(k)
			a = append(a, num)
		} else if v == maxCount {
			num, _ := strconv.Atoi(k)
			a = append(a, num)
		}
	}
	b := quickSort(a)
	return b
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := arr[0]
	var left []int
	var right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] > middle {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	middle_s := []int{middle}
	left = quickSort(left)
	right = quickSort(right)
	arr = append(append(left, middle_s...), right...)
	return arr
}
