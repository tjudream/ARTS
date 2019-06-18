package week13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func dfs(arr []int) int {
	sum := 0
	for i := 0; i < 26; i++ {
		if arr[i] == 0 {
			continue
		}
		arr[i]--
		sum = sum + 1 + dfs(arr)
		arr[i]++
	}
	return sum
}

func numTilePossibilities(tiles string) int {
	var arr []int = make([]int,26)
	for i := 0; i < len(tiles); i++ {
		arr[tiles[i] - 'A']++
	}
	return dfs(arr)
}

func TestNumTilePossiblities(t *testing.T) {
	e1 := "AAB"
	assert.Equal(t, 8, numTilePossibilities(e1))
	e2 := "AAABBC"
	assert.Equal(t, 188, numTilePossibilities(e2))
}
