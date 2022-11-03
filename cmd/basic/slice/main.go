package slice

import "fmt"

func main() {
	a := make([]int, 0, 10) //給足夠容量
	b := append(a, 1, 2, 3)
	_ = append(a, 99, 88)
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	// 因為底層是指針 指向
	//--------

	aa := make([]int, 0, 2) //給超小容量
	bb := append(aa, 1, 2, 3)
	_ = append(aa, 99, 88, 77)
	fmt.Println(bb)
}
