package longestPalindrome

import (
	"fmt"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	s := longestPalindrome2("babad")
	if strings.Compare(s, "bab") == 0 || strings.Compare(s, "aba") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test2(t *testing.T) {
	s := longestPalindrome2("cbbd")
	if strings.Compare(s, "bb") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test3(t *testing.T) {
	s := longestPalindrome2("aa")
	if strings.Compare(s, "aa") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test4(t *testing.T) {
	s := longestPalindrome2("aacabdkacaa")
	if strings.Compare(s, "aca") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test5(t *testing.T) {
	s := longestPalindrome2("aba")
	if strings.Compare(s, "aba") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test6(t *testing.T) {
	s := longestPalindrome2("a")
	if strings.Compare(s, "a") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func Test7(t *testing.T) {
	s := longestPalindrome2("aaaa")
	if strings.Compare(s, "aaaa") == 0 {
		t.Log("pass , s = ", s)
	} else {
		t.Error("fatal, s = ", s)
	}
}

func longestPalindrome(s string) string {
	var start, num, finalStart, maxNum int
	var c bool

	for n, _ := range s {
		if !c && finalStart == 0 && n > 0 {
			if s[n] == s[n-1] {
				finalStart = n
				c = true
				continue
			}
		}

		t := n - ((num + 1) * 2)
		if t < 0 {
			if num > maxNum {
				finalStart, maxNum = start, num
				c = false
			}
			start, num = 0, 0
			continue
		}

		if s[int(t)] == s[n] {
			num++
			if start == 0 {
				start = n
			}
			if num > maxNum {
				finalStart, maxNum = start, num
				c = false
			}
		} else {
			start, num = 0, 0
		}
	}

	fmt.Println("finalStart = ", finalStart, ", num = ", maxNum)
	if c {
		return s[finalStart-1 : finalStart+1]
	}
	if maxNum == 0 {
		return string(s[0])
	}
	if finalStart-maxNum-1 < 0 || finalStart+maxNum > len(s) {
		return ""
	}

	fmt.Println("end1 = ", finalStart-maxNum-1, " , end 2= ", finalStart+maxNum)
	return s[finalStart-maxNum-1 : finalStart+maxNum]
}

/*
71% ; 92%

這課學到的邏輯是有些地方:
1. 有些還是得要用兩回圈
2. 先釐清題目對於特別案例的要求結果 (釐清需求)
3. 測試案例盡量寫到 特別案例 ，像是符合的條件在頭尾，不同的組合
*/
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	var start, end, mn, ml, continuous int

	for i := 0; i < len(s); i++ {
		if i+2 < len(s) {
			if s[i] == s[i+2] {
				start, end = i, i+2
				continuous = 1
				if end-start > ml-mn {
					mn, ml = start, end
				}
				for {
					if i-continuous < 0 || i+2+continuous == len(s) {
						start, end, continuous = 0, 0, 0
						break
					}
					if s[i-continuous] == s[i+2+continuous] {
						start, end = i-continuous, i+2+continuous
						if end-start > ml-mn {
							mn, ml = start, end
						}
						continuous++
					} else {
						start, end, continuous = 0, 0, 0
						break
					}
				}
			}
		}
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				start, end = i, i+1
				continuous = 1
				if end-start > ml-mn {
					mn, ml = start, end
				}
				for {
					if i-continuous < 0 || i+1+continuous == len(s) {
						start, end, continuous = 0, 0, 0
						break
					}
					if s[i-continuous] == s[i+1+continuous] {
						start, end = i-continuous, i+1+continuous
						if end-start > ml-mn {
							mn, ml = start, end
						}
						continuous++
					} else {
						start, end, continuous = 0, 0, 0
						break
					}
				}
			}
		}
	}
	return s[mn : ml+1]
}
