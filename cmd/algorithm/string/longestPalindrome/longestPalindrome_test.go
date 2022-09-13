package longestPalindrome

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	s := longestPalindrome("babad")
	if strings.Compare(s, "bab") == 0 || strings.Compare(s, "aba") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test2(t *testing.T) {
	s := longestPalindrome("cbbd")
	if strings.Compare(s, "bb") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test3(t *testing.T) {
	s := longestPalindrome("aa")
	if strings.Compare(s, "aa") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test4(t *testing.T) {
	s := longestPalindrome("aacabdkacaa")
	if strings.Compare(s, "aca") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

//func Test5(t *testing.T) {
//	s := longestPalindrome("aa")
//	if strings.Compare(s, "aa") == 0 {
//		t.Log("pass , s = ", s)
//	} else {
//		t.Error("fatal, s = ", s)
//	}
//}

func longestPalindrome(s string) string {
	var start, to int

	for n, v := range s {
		first := strings.Index(s, string(v))
		if n+1-first > to-start {
			to, start = n+1, first
		}
	}

	return s[start:to]
}
