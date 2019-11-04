package week33

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minDistance(word1 string, word2 string) int {
	m,n := len(word1),len(word2)
	var dp [][]int
	dp = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				min := dp[i][j]
				if min > dp[i+1][j] {
					min = dp[i+1][j]
				}
				if min > dp[i][j+1] {
					min = dp[i][j+1]
				}
				dp[i+1][j+1] = min + 1
			}
		}
	}
	return dp[m][n]
}

func TestMinDistance(t *testing.T) {
	word1,word2 := "horse","ros"
	assert.Equal(t, minDistance(word1,word2),3)
	word1,word2 = "intention","execution"
	assert.Equal(t, minDistance(word1,word2),5)
}
