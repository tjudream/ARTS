package week25

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func preprocesses(s string) string {
	var sb strings.Builder
	sb.WriteRune('#')
	for _,e := range s {
		sb.WriteRune(e)
		sb.WriteRune('#')
	}
	return sb.String()
}

func longestPalindrome(s string) string {
	if len(s) <= 1{
		return s
	}
	s1 := preprocesses(s)
	len := len(s1)
	id,maxRight := 0,0
	p := make([]int, len)

	maxlen := -1
	resid := 0

	for i := 1; i < len; i++ {
		if i < maxRight {
			p[i] = p[2*id - i]
			if p[i] > maxRight - i {
				p[i] = maxRight - i
			}
		} else {
			p[i] = 1
		}

		for i - p[i] >=0 && i + p[i] < len && s1[i - p[i]] == s1[i + p[i]] {
			p[i]++
		}
		if maxRight < i + p[i] {
			id = i
			maxRight = i + p[i]
		}

		if maxlen < p[i] {
			maxlen = p[i]
			resid = id
		}
	}

	return strings.Replace(s1[resid - p[resid] + 1: resid + p[resid] - 1], "#", "", -1)
}

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	assert.Equal(t,"bab", longestPalindrome(s))
	s = "cbbd"
	assert.Equal(t, "bb", longestPalindrome(s))
}