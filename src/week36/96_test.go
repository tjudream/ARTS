package week36

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	var dp []int
	dp = make([]int,n+1)
	dp[0],dp[1] = 1,1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j]*dp[i-j-1]
		}
	}
	return dp[n]
}

func TestNumTrees(t *testing.T)  {
	assert.Equal(t, 5, numTrees(3));
}