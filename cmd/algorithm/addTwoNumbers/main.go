package addTwoNumbers

import (
	"fmt"
	"strconv"
)

func addTwoNumbers(list1, list2 []int) []int {
	for i, v := range list1 {
		list1[i] = v + list2[i]
		if list1[i] >= 10 {
			s := strconv.Itoa(list1[i])
			s = s[len(s)-1:]
			fmt.Println("s = ", s)
			atoi, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("err = ", err)
			}
			list1[i] = atoi
		}
	}
	return list1
}
