package secondTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"
	r := solution2(beginWord, endWord, wordList)

	assert.Equal(t, 5, r, "Test1 they should be equal, your answer is : %v", r)
}

func solution2(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	mp := make(map[string]bool)

	for _, w := range wordList {
		mp[w] = true
	}

	findInMap := func(str string) bool {
		_, ok := mp[str]
		return ok
	}

	queue := make([]string, 0)

	queue = append(queue, beginWord)

	steps := 0

	for len(queue) > 0 {
		for _, q := range queue {
			w := q
			if w == endWord {
				return steps + 1
			}
			for i := 0; i < len(w); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					ww := w[:i] + string(ch) + w[i+1:]
					if findInMap(ww) {
						queue = append(queue, ww)
						delete(mp, ww)
					}
				}
			}
			queue = queue[1:]
		}
		steps++
	}

	return 0
}
