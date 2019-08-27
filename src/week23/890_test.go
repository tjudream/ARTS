package week23

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	lp := len(pattern)
	for _,word := range words {
		var m = make(map[uint8]uint8)
		var n = make(map[uint8]uint8)
		l := len(word)
		if  l != lp {
			continue
		}
		flg := true
		for i := 0; i < l; i++ {
			if m[pattern[i]] == 0 && n[word[i]] == 0 {
				m[pattern[i]] = word[i]
				n[word[i]] = pattern[i]
			} else if m[pattern[i]] == word[i] && n[word[i]] == pattern[i] {
				continue
			} else {
				flg = false
				break
			}
		}
		if flg {
			res = append(res, word)
		}
	}
	return res
}

func TestFindAndReplacePattern(t *testing.T)  {
	words := []string{"abc","deq","mee","aqq","dkd","ccc"}
	res := []string{"mee","aqq"}
	pattern := "abb"
	assert.Equal(t, res, findAndReplacePattern(words, pattern))
}
