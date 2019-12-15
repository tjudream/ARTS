package week38

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isInter(s1 string, s2 string, s3 string) bool {
	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}
	return (s3[0] == s1[0] && isInter(s1[1:len(s1)], s2, s3[1:len(s3)])) || (s3[0] == s2[0] && isInter(s1, s2[1:len(s2)], s3[1:len(s3)]))
}
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l3 != l1+l2 {
		return false
	}
	return isInter(s1, s2, s3)
}
func isInterleaveDp(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l3 != l1+l2 {
		return false
	}
	var dp [][]bool
	dp = make([][]bool, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]bool, l2+1)
	}
	dp[0][0] = true
	for i := 1; i < l1+1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j < l2+1; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[l1][l2]
}

func TestIsInterleave(t *testing.T) {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbcbcac"
	assert.True(t, isInterleave(s1, s2, s3))
	assert.True(t, isInterleaveDp(s1, s2, s3))
	s1, s2, s3 = "aabcc", "dbbca", "aadbbbaccc"
	assert.False(t, isInterleave(s1, s2, s3))
	assert.False(t, isInterleaveDp(s1, s2, s3))
	s1, s2, s3 = "", "", ""
	assert.True(t, isInterleave(s1, s2, s3))
	assert.True(t, isInterleaveDp(s1, s2, s3))
}
