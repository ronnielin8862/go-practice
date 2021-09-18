package testifyMathForTest

import (
	"fmt"
	"testing"

	"github.com/ronnielin8862/go-api/pkg/mathForTest"
	"github.com/stretchr/testify/assert"
)

func TestMathPlus(t *testing.T) {
	fmt.Println("開始測試 加法")

	assert := assert.New(t)

	var a, b = 2, 3

	assert.NotNil(mathForTest.Plus(a, b), "不應該是nil")

	//測試Error
	// assert.Equal(mathForTest.Plus(a, b), mathForTest.Minus(a, b), "應該要相等")

	t.Logf("又出事了")
}
