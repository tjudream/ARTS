package week24

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestPalindrome(s string) string {
	l,start,end := len(s),0,0
	if l <= 1 {
		return s
	}
	for i := 0; i < l; i++ {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i + 1)
		len := l1
		if l2 > len {
			len = l2
		}
		if len > (end - start) {
			start = i - (len - 1)/2
			end = i + len/2
		}
	}
	return s[start:end + 1]
}

func expandAroundCenter(s string, left int, right int) int {
	l := len(s)
	for left >= 0 && right < l && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	assert.Equal(t,"aba", longestPalindrome(s))
	s = "cbbd"
	assert.Equal(t, "bb", longestPalindrome(s))
}
