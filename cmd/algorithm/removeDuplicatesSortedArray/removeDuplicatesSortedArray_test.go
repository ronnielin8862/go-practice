package removeDuplicatesSortedArray

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	numbers := []int{1, 1, 2}
	wantNum := 2
	wantNewNums := []int{1, 2}

	num, newNums := removeDuplicates(numbers)
	check(t, num, newNums, wantNum, wantNewNums)
}

func Test2(t *testing.T) {
	numbers := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	wantNum := 5
	wantNewNums := []int{0, 1, 2, 3, 4}

	num, newNums := removeDuplicates(numbers)
	check(t, num, newNums, wantNum, wantNewNums)
}

func Test3(t *testing.T) {
	numbers := []int{1}
	wantNum := 1
	wantNewNums := []int{1}

	num, newNums := removeDuplicates(numbers)
	check(t, num, newNums, wantNum, wantNewNums)
}

func check(t *testing.T, num int, newNums []int, wantNum int, wantNewNums []int) {
	fmt.Print("newNums: ", newNums, " num: ", num)

	if num != wantNum {
		t.Errorf("got %v wantNum %v given", num, wantNum)
	}
	for i := 0; i < num; i++ {
		if newNums[i] != wantNewNums[i] {
			t.Errorf("got %v wantNum %v given", newNums[i], wantNewNums[i])
		}
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

func removeDuplicates(numbers []int) (num int, newNums []int) {
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
	return
}
