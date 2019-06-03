package week11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func sortArrayByParity(A []int) []int {
	s := 0
	e := len(A) - 1
	for s < e {
		for A[s]&1 == 0 && s < e {
			s++
		}
		for A[e]&1 == 1 && s < e {
			e--
		}
		if (s < e) {
			A[s],A[e] = A[e],A[s]
		}
	}
	return A
}

func TestSortArrayByParity(t *testing.T) {
	A := []int{3,1,2,4}
	assert.Equal(t, []int{4,2,1,3}, sortArrayByParity(A))
}
