package week12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		leni := len(A[i]);
		for j := 0; j < (leni + 1)/2; j++ {
			A[i][j],A[i][leni - j - 1] = A[i][leni - j - 1]^1,A[i][j]^1
		}
	}
	return A
}

func TestFlipAndInvertImage(t *testing.T) {
	A1 := [][]int{{1,1,0},{1,0,1},{0,0,0}}
	E1 := [][]int{{1,0,0},{0,1,0},{1,1,1}}
	assert.Equal(t, E1, flipAndInvertImage(A1))
	A2 := [][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}}
	E2 := [][]int{{1,1,0,0},{0,1,1,0},{0,0,0,1},{1,0,1,0}}
	assert.Equal(t, E2, flipAndInvertImage(A2))
}

