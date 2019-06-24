package week14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Time Complex: O(N); Space Complex: O(N)
func repeatedNTimes(A []int) int {
	dict := make([]int, 10000)
	for i := 0; i < len(A); i++ {
		if dict[A[i]] == 0 {
			dict[A[i]] = 1
		} else {
			return A[i]
		}
	}
	return -1
}

// Time Complex: O(N); Space Complex: O(1)
func repeatedNTimes2(A []int) int {
	a := -1
	b := -1
	c := -1
	for i := 0; i < len(A); i++ {
		if A[i] == a {
			return a
		}
		if A[i] == b {
			return b
		}
		if A[i] == c {
			return c
		}
		a = b
		b = c
		c = A[i]
	}
	return -1
}

func TestRepeatedNTimes(t *testing.T) {
	A := []int{1,2,3,3}
	assert.Equal(t, 3, repeatedNTimes(A))
	A = []int{2,1,2,5,3,2}
	assert.Equal(t, 2, repeatedNTimes(A))
	A = []int{5,1,5,2,5,3,5,4}
	assert.Equal(t, 5, repeatedNTimes(A))
}

func TestRepeatedNTimes2(t *testing.T) {
	A := []int{1,2,3,3}
	assert.Equal(t, 3, repeatedNTimes2(A))
	A = []int{2,1,2,5,3,2}
	assert.Equal(t, 2, repeatedNTimes2(A))
	A = []int{5,1,5,2,5,3,5,4}
	assert.Equal(t, 5, repeatedNTimes2(A))
}