package mathForTest

import (
	"fmt"
	"testing"

	"github.com/ronnielin8862/go-api/pkg/mathForTest"
)

func TestMathPlus(t *testing.T) {
	fmt.Println("開始測試 加法")

	var a, b = 2, 3

	result := mathForTest.Plus(a, b)

	if result != a+b {
		t.Error("加法錯誤")
	}

	t.Logf("又錯了，還不好好寫程式")
}

func TestMathMinus(t *testing.T) {
	fmt.Println("開始測試 減法")

	var a, b = 5, 1

	result := mathForTest.Minus(a, b)

	if result != a-b {
		t.Error("減法錯誤")
	}

	t.Logf("又錯了，還不好好寫程式")
}
