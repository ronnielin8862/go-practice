package diverseString

import (
	"fmt"
	"testing"
)

func TestLoopString(t *testing.T) {
	stringArray := diverseString("Ronnie")
	fmt.Println(len(stringArray) == 2*2*2*2*2)

	for _, c := range stringArray {
		fmt.Println("list = ", c)
	}
}

func TestAdmin(t *testing.T) {
	stringArray := diverseString("aba")
	m := map[string]struct{}{
		"aba": {},
		"abA": {},
		"aBa": {},
		"aBA": {},
		"Aba": {},
		"AbA": {},
		"ABa": {},
		"ABA": {},
	}

	for _, c := range stringArray {
		fmt.Println("list = ", c)
	}

	for k, _ := range m {
		has := false
		for _, c := range stringArray {
			//fmt.Printf("k = %s , c = %s \n", k, c)
			if c == k {
				has = true
				break
			}
		}
		if !has {
			t.Error("answer not including ", k)
		}
	}
}

func TestGeneratePermutations(t *testing.T) {
	generatePermutations("aba")
}

func TestLoopString2(t *testing.T) {
	generatePermutations("Ronnie")
}
