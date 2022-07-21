package main

import (
	"fmt"
	"strings"
)

func isIsomorphic(s string, t string) bool {

	for i := 0; i < len(s); i++ {

		var sameS []int8
		var sameT []int8

		for n, v := range s {
			var targetS = int8(s[i])
			if int8(v) == targetS {
				sameS = append(sameS, int8(n))
			}
		}
		for n, v := range t {
			var targetT = int8(t[i])
			if int8(v) == targetT {
				sameT = append(sameT, int8(n))
			}
		}

		if len(sameS) != len(sameT) {
			return false
		}
		for k, v := range sameS {
			if v != sameT[k] {
				return false
			}
		}
	}

	return true
}

func isIsomorphic2(s string, t string) bool {
	m := make(map[string]string)

	for i := 0; i < len(s); i++ {
		if ok := m[string(s[i])]; ok == "" {
			if strings.Contains(t[:i], string(t[i])) {
				return false
			}
			m[string(s[i])] = string(t[i])
		} else {
			if ok != string(t[i]) {
				return false
			}
		}
	}

	return true
}

func main() {

	//s := "egg"
	//t := "add"
	//fmt.Println(isIsomorphic(s, t))
	//
	//a := "foo"
	//b := "bar"
	//fmt.Println(isIsomorphic(a, b))
	//
	//x := "paper"
	//y := "title"
	//fmt.Println(isIsomorphic(x, y))
	//
	//h := "bbbaaaba"
	//j := "aaabbbba"
	//fmt.Println(isIsomorphic(h, j))

	rr := "badc"
	ii := "baba"
	fmt.Println(isIsomorphic2(rr, ii))

	ss := "egg"
	tt := "add"
	fmt.Println(isIsomorphic2(ss, tt))

	aa := "foo"
	bb := "bar"
	fmt.Println(isIsomorphic2(aa, bb))

	xx := "paper"
	yy := "title"
	fmt.Println(isIsomorphic2(xx, yy))

	hh := "bbbaaaba"
	jj := "aaabbbba"
	fmt.Println(isIsomorphic2(hh, jj))

}
