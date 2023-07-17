package removeDuplicatesSortedArray

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	numbers := []int{1, 1, 2}
	wantNum := 2

	num := removeDuplicates(numbers)
	check(t, num, wantNum)
}

func Test2(t *testing.T) {
	numbers := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	wantNum := 5

	num := removeDuplicates(numbers)
	check(t, num, wantNum)
}

func Test3(t *testing.T) {
	numbers := []int{1}
	wantNum := 1
	num := removeDuplicates(numbers)

	check(t, num, wantNum)
}

func Test4(t *testing.T) {
	numbers := []int{1, 1}
	wantNum := 1
	num := removeDuplicates(numbers)

	check(t, num, wantNum)
}

func check(t *testing.T, num int, wantNum int) {
	if num != wantNum {
		t.Errorf("got %v wantNum %v given", num, wantNum)
	}
}

func arrayContains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func removeDuplicates(numbers []int) (num int) {
	var newNums []int
	if len(numbers) == 0 {
		return
	}
	newNums = append(newNums, numbers[0])
	num = 1
	for i := 1; i < len(numbers); i++ {
		if !arrayContains(newNums, numbers[i]) {
			newNums = append(newNums, numbers[i])
			num++
		}
	}
	fmt.Println(newNums, num)
	return
}
