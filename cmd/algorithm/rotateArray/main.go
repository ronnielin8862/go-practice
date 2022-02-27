package main

import "fmt"

//给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。
//
//示例 1:
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//示例2:
//
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100]

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	//method1(nums, k)
	//method2(nums, k)
	method3(nums, k)
}

//T(n) = O(n^2); S(n) = O(n)
func method1(nums []int, k int) {
	for i := 1; i <= k; i++ {

		var newNums []int
		lastNum := nums[len(nums)-1]
		newNums = append(newNums, lastNum)

		for j := 0; j < len(nums)-1; j++ {
			newNums = append(newNums, nums[j])
		}

		copy(nums, newNums)
		fmt.Println(nums)
	}
}

//T(n) = O(n); S(n) = O(1)
func method2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	fmt.Println(nums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

//T(n) = O(n); S(n) = O(n)
func method3(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
	fmt.Println(nums)
}
