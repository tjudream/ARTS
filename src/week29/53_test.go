package week29

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxSubArray(nums []int) int {
	lenn := len(nums)
	if lenn == 0 {
		return 0
	}
	var dp []int
	dp = make([]int, lenn)
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < lenn; i++ {
		dp[i] = dp[i-1] + nums[i]
		if dp[i] < nums[i] {
			dp[i] = nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum := 6
	assert.Equal(t, sum, maxSubArray(nums))
}
