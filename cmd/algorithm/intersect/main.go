package main

import "fmt"

//给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//示例 2:
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]

func main() {
	//nums1 := []int{4, 9, 5}
	//nums2 := []int{9, 4, 9, 8, 4}

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	//newNums := method1(nums1, nums2)

	newNums := method2(nums1, nums2)

	fmt.Println(newNums)
}

// T = O(m+n + m^2); S = O(m+n+ min(m,n))
func method1(nums1 []int, nums2 []int) []int {
	numMap1 := make(map[int]int)
	numMap2 := make(map[int]int)
	var arr []int

	for _, v := range nums1 {
		numMap1[v]++
	}

	for _, v := range nums2 {
		numMap2[v]++
	}

	for k, map1Value := range numMap1 {
		map2Value := numMap2[k]

		if map2Value != 0 {
			if map2Value > map1Value {
				for i := 1; i <= map1Value; i++ {
					arr = append(arr, k)
				}
			} else if map2Value <= map1Value {
				for i := 1; i <= map2Value; i++ {
					arr = append(arr, k)
				}
			}
		}
	}
	return arr
}

// T = O(m+n); S = O(min(m,n))
func method2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return method2(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}
