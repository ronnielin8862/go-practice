package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if !strings.Contains(t, string(s[i])) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
	fmt.Println(isSubsequence("acb", "ahbgdc")) // false
}
