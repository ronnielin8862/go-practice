package main

import "fmt"

func runningSum(nums []int) []int {
	var newNums []int
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		newNums = append(newNums, total)
	}
	return newNums
}

func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
	fmt.Println(runningSum2(nums))
}
