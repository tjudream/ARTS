package week27

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestValidParentheses(s string) int {
	left,right,max,lens := 0,0,0,len(s)

	for i := 0; i < lens; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			newMax := 2 * left
			if newMax > max {
				max = newMax
			}
		} else if right >= left {
			left,right = 0,0
		}
	}

	left,right = 0,0
	for i := lens - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			newMax := 2*left
			if newMax > max {
				max = newMax
			}
		} else if left >= right {
			left,right = 0,0
		}
	}
	return max
}

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	assert.Equal(t,2,longestValidParentheses(s))
	s = ")()())"
	assert.Equal(t,4,longestValidParentheses(s))
}
