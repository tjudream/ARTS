package week28

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	lens,lenp := len(s),len(p)
	var dp [][]bool
	dp = make([][]bool, lens + 1)
	for i := 0; i < lens + 1; i++ {
		dp[i] = make([]bool, lenp + 1)
	}

	dp[0][0] = true
	for j := 1; j <= lenp; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i <= lens; i++ {
		for j := 1; j <= lenp; j++ {
			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[lens][lenp]
}

func TestIsMatch(t *testing.T) {
	s := "aa"
	p := "a"
	assert.False(t, isMatch(s, p))
	s = "aa"
	p = "*"
	assert.True(t, isMatch(s, p))
	s = "cb"
	p = "?a"
	assert.False(t, isMatch(s, p))
	s = "adceb"
	p = "*a*b"
	assert.True(t, isMatch(s, p))
	s = "acdcb"
	p = "a*c?b"
	assert.False(t, isMatch(s, p))
}