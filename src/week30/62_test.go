package week30
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func uniquePaths(m int, n int) int {
    var dp [][]int
    dp = make([][]int, m)
    for i := 0; i < m; i++ {
    	dp[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}

func TestUniquePaths(t *testing.T) {
    m,n := 3,2
    assert.Equal(t, 3, uniquePaths(m,n))
    m,n = 7,3
    assert.Equal(t, 28, uniquePaths(m,n))
}
