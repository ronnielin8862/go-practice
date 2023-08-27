package RemoveDuplicatesSortedArrayII

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}      // Input array
	expectedNums := []int{1, 1, 2, 2, 3} // The expected answer with correct length

	expectLen := removeDuplicates(nums) // Calls your implementation

	if expectLen != len(expectedNums) {
		t.Fatalf("Wrong result for %v", nums)
	}

	for i := 0; i < expectLen; i++ {
		if expectedNums[i] != nums[i] {
			t.Fatalf("Wrong result for %v", nums)
		}
	}
}

func Test2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}   // Input array
	expectedNums := []int{0, 0, 1, 1, 2, 3, 3} // The expected answer with correct length

	expectLen := removeDuplicates(nums) // Calls your implementation

	if expectLen != len(expectedNums) {
		t.Fatalf("Wrong result for %v", nums)
	}

	for i := 0; i < expectLen; i++ {
		if expectedNums[i] != nums[i] {
			t.Fatalf("Wrong result for %v", nums)
		}
	}
}

func Test3(t *testing.T) {
	nums := []int{1}         // Input array
	expectedNums := []int{1} // The expected answer with correct length

	expectLen := removeDuplicates(nums) // Calls your implementation

	if expectLen != len(expectedNums) {
		t.Fatalf("Wrong result for %v", nums)
	}

	for i := 0; i < expectLen; i++ {
		if expectedNums[i] != nums[i] {
			t.Fatalf("Wrong result for %v", nums)
		}
	}
}

func Test4(t *testing.T) {
	nums := []int{1, 2, 2, 2}           // Input array
	expectedNums := []int{1, 2, 2}      // The expected answer with correct length
	expectLen := removeDuplicates(nums) // Calls your implementation

	if expectLen != len(expectedNums) {
		t.Fatalf("Wrong result for %v", nums)
	}

	for i := 0; i < expectLen; i++ {
		if expectedNums[i] != nums[i] {
			t.Fatalf("Wrong result for %v", nums)
		}
	}
}

func Test5(t *testing.T) {
	nums := []int{1, 1}                 // Input array
	expectedNums := []int{1, 1}         // The expected answer with correct length
	expectLen := removeDuplicates(nums) // Calls your implementation

	if expectLen != len(expectedNums) {
		t.Fatalf("Wrong result for %v", nums)
	}

	for i := 0; i < expectLen; i++ {
		if expectedNums[i] != nums[i] {
			t.Fatalf("Wrong result for %v", nums)
		}
	}
}

func Test6(t *testing.T) {
	nums := []int{1, 1, 1}      // Input array
	expectedNums := []int{1, 1} // The expected answer with correct length

	expectLen := removeDuplicates(nums) // Calls your implementation

	if expectLen != len(expectedNums) {
		t.Fatalf("Wrong result for %v", nums)
	}

	for i := 0; i < expectLen; i++ {
		if expectedNums[i] != nums[i] {
			t.Fatalf("Wrong result for %v", nums)
		}
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	var skip = 0

	for i, t := range nums {
		if i <= 1 {
			continue
		}

		if skip > 0 {
			nums[i-skip] = t
			nums[i] = 0
		}

		if nums[i-skip-1] == t && nums[i-skip-2] == t {
			nums[i] = 0
			skip++
		}
	}

	return len(nums) - skip
}

func TestNil(t *testing.T) {
	nums := []int{3, 3, 3}
	for i, t := range nums {
		fmt.Println("i = ", i)
		fmt.Println("t = ", t)
	}
}
