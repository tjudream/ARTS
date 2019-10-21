package week31
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    m,n := len(obstacleGrid), len(obstacleGrid[0])
    var dp [][]int
    dp = make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = 1
    for i := 1; i < n; i++ {
        if obstacleGrid[0][i] == 1 {
            break
        }
        dp[0][i] = 1
    }
    for i := 1; i < m; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
                continue
            }
            if obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
                dp[i][j] = 0
            } else if obstacleGrid[i-1][j] == 1 {
                dp[i][j] = dp[i][j-1]
            } else if obstacleGrid[i][j-1] == 1 {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i][j-1] + dp[i-1][j]
            }
        }
    }
    return dp[m-1][n-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
    obstacleGrid := [][]int{{0,0,0},{0,1,0},{0,0,0}}
    assert.Equal(t, 2, uniquePathsWithObstacles(obstacleGrid))
}