package main

import "fmt"

//https: //leetcode-cn.com/problems/add-digits/
//给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
//示例 1:
//
//输入: num = 38
//输出: 2
//解释: 各位相加的过程为：
//38 --> 3 + 8 --> 11
//11 --> 1 + 1 --> 2
//由于2 是一位数，所以返回 2。
//示例 1:
//
//输入: num = 0
//输出: 0

func main() {
	num := 388
	fmt.Println(addDigits(num))
}

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for ; num > 0; num /= 10 {
			sum += num % 10
		}
		num = sum
	}
	return num
}

//
//func addDigits(num int) int {
//	var t int
//	var r int
//	numString := strconv.Itoa(num)
//	fmt.Println("numString = ", numString)
//	for i, _ := range numString {
//		fmt.Println(numString[i])
//		//t += int(v)
//		fmt.Println(t)
//		break
//	}
//	for i := 0; i < len(numString); i++ {
//		fmt.Println(numString)
//		fmt.Println("ddd = ", numString[i])
//	}
//	//if t >= 10 {
//	//	r = addDigits(t)
//	//}
//	return r
//}
