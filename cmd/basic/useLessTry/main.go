package main

import "fmt"

func main() {
	var ints []int

	ints = append(ints, 1)
	ints = append(ints, 2)
	ints = append(ints, 3)
	ints = append(ints, 4)

	//fmt.Println(ints)
	//for _, v := range ints {
	//	fmt.Println(v)
	//}
	//

	s := len(ints)

	for i := 0; i < s; i++ {
		ns := s - i - 1

		fmt.Println(ints[ns])
	}

}
