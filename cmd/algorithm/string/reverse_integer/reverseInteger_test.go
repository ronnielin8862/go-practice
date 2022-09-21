package reverse_integer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/880/
// 字符串移位需轉乘其他類型，例如[]byte
// todo 未完成
func TestReverseInteger(t *testing.T) {
	var a int32
	var b int32
	//fmt.Println(reverse(123))
	//fmt.Println(reverse(-123))
	//fmt.Println(reverse(120))
	//fmt.Println(reverse(1534236469))
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)

	fmt.Println(a)
	fmt.Println(b)

	c := reverse(123)
	if c != 321 {
		t.Fatalf("reverse(123) = %d; want 321", c)
	}

	c = reverse(-123)
	if c != -321 {
		t.Fatalf("reverse(-123) = %d; want -321", c)
	}

	c = reverse(120)
	if c != 21 {
		t.Fatalf("reverse(120) = %d; want 21", c)
	}

	c = reverse(1534236469)
	if c != 0 {
		t.Fatalf("reverse(1534236469) = %d; want 0", c)
	}
}

func reverse(n int32) int32 {
	// 確認是否為負值，記號下來
	isNegtive := false
	if n < 0 {
		isNegtive = true
	}

	// 反轉數字
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	b := bytesBuffer.Bytes()
	for i := 0; i < len(b); i++ {
		j := len(b) - 1

		b[i], b[j] = b[j], b[i]
	}

	r := bytes.NewBuffer(b)
	n = 0
	binary.Read(r, binary.BigEndian, n)
	fmt.Println("n1 = ", n)
	// 確認是否超過32位元範圍
	// 負值加上負號
	if isNegtive {
		n = n * -1
	}

	return n
}
