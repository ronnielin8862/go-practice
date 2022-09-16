package byte

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	case1 := "hello"
	assert.Equal(t, "olleh", reverseString([]byte(case1)))

	case2 := "Hannah"
	assert.Equal(t, "hannaH", reverseString([]byte(case2)))
	//case2 := []string{"H", "a", "n", "n", "a", "h"}
}

func reverseString(s []byte) string {
	a := string(s)
	l := len(a)
	for i := 0; i < l/2; i++ {

		if i < 1 {
			a = string(a[(l-1)]) + a[1:l-1] + string(a[0])
		} else {
			a = a[0:i] + string(a[(l-1)]) + a[i+1:l-1] + string(a[i]) + a[l:]
		}
		l--
	}

	return a
}
