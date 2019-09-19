package week26

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

	dp[lens][lenp] = true
	for i := lens; i >= 0; i-- {
	    for j := lenp - 1; j >= 0; j-- {
	        first_match := i < lens &&  (p[j] == s[i] || p[j] == '.')
	        if (j + 1 < lenp && p[j+1] == '*') {
				dp[i][j] =  dp[i][j+2] || first_match && dp[i+1][j];
	        } else {
	            dp[i][j] = first_match && dp[i+1][j+1];
	        }
	    }
	}
	return dp[0][0]
}

func TestIsMatch(t *testing.T) {
	s,p := "aa","a"
	assert.False(t,isMatch(s, p))
	s,p = "aa","a*"
	assert.True(t, isMatch(s, p))
	s,p = "ab",".*"
	assert.True(t, isMatch(s, p))
	s,p = "aab","c*a*b*"
	assert.True(t, isMatch(s, p))
	s,p = "mississippi","mis*is*p*."
	assert.False(t, isMatch(s, p))
	s,p = "",""
	assert.True(t, isMatch(s, p))
	s,p = "",".*"
	assert.True(t, isMatch(s,p))
	s,p = "a",""
	assert.False(t, isMatch(s, p))
	s,p = "abc","abc"
	assert.True(t, isMatch(s, p))
}