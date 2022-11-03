package slice

import (
	"fmt"
	//"github.com/influxdata/influxdb/pkg/slices"
	"golang.org/x/exp/slices"
	"testing"
)

type (
	User struct {
		Id  int `json:"id"`
		Age int `json:"age"`
		Adder
		Experience Experience `json:"experience"`
	}

	Adder struct {
		Street string `json:"street"`
		Number int    `json:"number"`
		Phone  `json:"phone"`
	}

	Experience struct {
		Company string `json:"company"`
		Job     string `json:"job"`
	}

	Phone struct {
		HomeNum string `json:"home_num"`
	}
)

// 這裡測試結果 = userA != userB。
// 說明了雙等號 可以比較到表層，以及直接引用成為的表層struct
func Test_struct_Compare1(t *testing.T) {
	userA := User{Id: 1, Age: 18, Adder: Adder{Street: "streetAAA", Number: 1}}
	userB := User{Id: 1, Age: 18, Adder: Adder{Street: "streetBBB", Number: 2}}

	if userA == userB {
		print("userA == userB")
	} else {
		print("userA != userB")
	}
}

// 這裡測試結果 = userA != userB。
// 說明了雙等號 可以比較到表層，以及第一層子結構
func Test_struct_Compare2(t *testing.T) {
	userA := User{Id: 1, Age: 18, Experience: Experience{Company: "companyA", Job: "jobA"}}
	userB := User{Id: 1, Age: 18, Experience: Experience{Company: "companyB", Job: "jobB"}}

	if userA == userB {
		print("userA == userB")
	} else {
		print("userA != userB")
	}
}

func Test_struct_Compare3(t *testing.T) {
	userA := User{Id: 1, Age: 18, Adder: Adder{Street: "streetAAA", Number: 1, Phone: Phone{HomeNum: "123"}}}
	userB := User{Id: 1, Age: 18, Adder: Adder{Street: "streetAAA", Number: 1, Phone: Phone{HomeNum: "456"}}}

	if userA == userB {
		print("userA == userB")
	} else {
		print("userA != userB")
	}
}

// 以上至此的測試結論是 : struct  可以直接比較，會進行深度的比較。    但是！  如果結構內包含指針，那比較結果肯定會返回不一致； 再來，如果結構內還有不可比較的類型: slice, map , func 時，比較就會報錯
// 面對需要比較slice 的結構體，可以使用 reflect.DeepEqual 進行比較

// 如果只是兩個slice 比較，可以使用 slices.Compare

func TestSlicesCompare(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 2}
	c := []string{"王小明", "王小華", "王小張"}
	d := []string{"王小明", "王小華", "王小髒"}
	//print(slices.e(c, "王小華"))

	//if a == b{
	//
	//}
	print("compare ? ", slices.Compare(a, b))

	print("equal ? ", slices.Equal(a, b))
	print("\n")

	fmt.Println("compare ? ", slices.Compare(c, d))
	fmt.Println("equal ? ", slices.Equal(c, d))

	fmt.Println("移除 元素王小華 ", slices.Delete(c, 1, 2))
	fmt.Println("移除 元素王小華 以後原本的 c 會被改變嗎", c) // 非指針類型不會被改變

}
