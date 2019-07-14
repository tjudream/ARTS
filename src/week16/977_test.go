package week16

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i]*A[i]
	}
	sort.Ints(A)
	return A
}

func sortedSquaresON(A []int) []int {
	N := len(A)
	min,minpos := math.MaxInt32,0
	for i := 0; i < N; i++ {
		A[i] = A[i]*A[i]
		if min > A[i] {
			min = A[i]
			minpos = i
		}
	}
	res := []int{A[minpos]}
	low := minpos - 1
	large := minpos + 1
	for low >= 0 && large < N {
		if A[low] < A[large] {
			res = append(res, A[low])
			low--
		} else {
			res = append(res, A[large])
			large++
		}
	}
	for low >= 0 {
		res = append(res, A[low])
		low--
	}
	if large < N {
		res = append(res, A[large:]...)
	}
	return res
}

func TestSortedSquares(t *testing.T) {
	I1 := []int{-4,-1,0,3,10}
	O1 := []int{0,1,9,16,100}
	assert.Equal(t, O1, sortedSquares(I1))
	I2 := []int{-7,-3,2,3,11}
	O2 := []int{4,9,9,49,121}
	assert.Equal(t, O2, sortedSquares(I2))
}

func TestSortedSquaresON(t *testing.T) {
	I1 := []int{-4,-1,0,3,10}
	O1 := []int{0,1,9,16,100}
	assert.Equal(t, O1, sortedSquaresON(I1))
	I2 := []int{-7,-3,2,3,11}
	O2 := []int{4,9,9,49,121}
	assert.Equal(t, O2, sortedSquaresON(I2))
	I3 := []int{-2,0}
	O3 := []int{0,4}
	assert.Equal(t, O3, sortedSquaresON(I3))
	I4 := []int{-3,-1,0}
	O4 := []int{0,1,9}
	assert.Equal(t, O4, sortedSquaresON(I4))
}
