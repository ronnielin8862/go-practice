package lengthOfLongestSubstring

import (
	"fmt"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	n := lengthOfLongestSubstring4("abcabcbb")
	//t.Log(n)
	if n == 3 {
		t.Log("1 success")
	} else {
		t.Error("1 fail")
	}
}

func Test2(t *testing.T) {
	n := lengthOfLongestSubstring4("pwwkew")
	//t.Log(n)
	if n == 3 {
		t.Log("2 success")
	} else {
		t.Error("2 fail")
	}
}

func Test3(t *testing.T) {
	n := lengthOfLongestSubstring4("wwkewkeio")
	//t.Log(n)
	if n == 5 {
		t.Log("3 success")
	} else {
		t.Error("3 fail, n = ", n)
	}
}

func Test4(t *testing.T) {
	n := lengthOfLongestSubstring4("ababcdeabce")
	//t.Log(n)
	if n == 5 {
		t.Log("4 success")
	} else {
		t.Error("4 fail")
	}
}

func Test5(t *testing.T) {
	n := lengthOfLongestSubstring4(" ")
	//t.Log(n)
	if n == 1 {
		t.Log("5 success")
	} else {
		t.Error("5 fail")
	}
}

func Test6(t *testing.T) {
	n := lengthOfLongestSubstring4("au")
	if n == 2 {
		t.Log("6 success")
	} else {
		t.Error("6 fail")
	}
}
func Test7(t *testing.T) {
	n := lengthOfLongestSubstring4("aa")
	if n == 1 {
		t.Log("7 success")
	} else {
		t.Error("7 fail")
	}
}

func Test8(t *testing.T) {
	n := lengthOfLongestSubstring4("dvdf")
	if n == 3 {
		t.Log("8 success")
	} else {
		t.Error("8 fail, n = ", n)
	}
}

// 基礎解法
func lengthOfLongestSubstring(s string) int {

	max := make(map[string]struct{})
	for i := 0; i < len(s); i++ {

		m := make(map[string]struct{})
		for j := i; j < len(s); j++ {

			if _, ok := m[string(s[j])]; ok {
				if len(m) > len(max) {
					max = m
				}
				break
			}
			m[string(s[j])] = struct{}{}

			if j == len(s)-1 && len(m) > len(max) {
				max = m
			}
		}
	}
	fmt.Println(max)
	return len(max)
}

// 空間複雜度 O(n), 時間複雜度 O(n)^2
func lengthOfLongestSubstring2(s string) int {
	var max string
	if len(s) > 0 {
		max = string(s[0])
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if strings.Contains(s[i:j], string(s[j])) {
				if j-i > len(max) {
					max = s[i:j]
				}
				break
			}

			// 最後一個
			if j+1 == len(s) && j+1-i > len(max) {
				max = s[i : j+1]
			}
		}
	}
	fmt.Println("s= ", max)
	return len(max)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 網路解 最快，不知道為何
func lengthOfLongestSubstring3(s string) int {
	maxLen, start := 0, 0
	table := [128]int{}
	for i, _ := range table {
		table[i] = -1
	}

	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		maxLen = maxInt(maxLen, i-start+1)
	}
	return maxLen
}

// 最佳解 ！  84% ; 100%  學起起來這種邏輯。  要先克服直覺想到的雙迴圈解法
func lengthOfLongestSubstring4(s string) int {
	start := 0
	var maxLength int

	for i := 0; i < len(s); i++ {
		if !strings.Contains(s[start:i], string(s[i])) {
			maxLength = maxInt(maxLength, i-start+1)
			if maxLength < i-start+1 {
				maxLength = i - start + 1
			}
		} else {
			start = strings.LastIndex(s[:i], string(s[i])) + 1
		}
	}
	return maxLength
}
