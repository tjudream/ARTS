package week32
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func minPathSum(grid [][]int) int {
    m,n := len(grid),len(grid[0])
    var s [][]int
    s = make([][]int,m)
    for i := 0; i < m; i++ {
        s[i] = make([]int,n)
    }

    s[0][0] = grid[0][0]

    for i := 1; i < m; i++ {
        s[i][0] = s[i-1][0] + grid[i][0]
    }
    for j := 1; j < n; j++ {
        s[0][j] = s[0][j-1] + grid[0][j]
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            min := s[i-1][j]
            if s[i][j-1] < s[i-1][j] {
                min = s[i][j-1]
            }
            s[i][j] = min + grid[i][j]
        }
    }
    return s[m-1][n-1]
}

func minPathSumO1Space(grid [][]int) int {
    m,n := len(grid),len(grid[0])
    for i := 1; i < m; i++ {
        grid[i][0] = grid[i-1][0] + grid[i][0]
    }
    for j := 1; j < n; j++ {
        grid[0][j] = grid[0][j-1] + grid[0][j]
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            min := grid[i-1][j]
            if grid[i][j-1] < grid[i-1][j] {
                min = grid[i][j-1]
            }
            grid[i][j] = min + grid[i][j]
        }
    }
    return grid[m-1][n-1]
}

func TestMinPathSum(t *testing.T) {
    grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
    assert.Equal(t, 7, minPathSum(grid))
    assert.Equal(t, 7, minPathSumO1Space(grid))
}