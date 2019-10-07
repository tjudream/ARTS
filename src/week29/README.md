# week29

---

# Algorithm [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)
## 1. 问题描述
子数组最大和

给定一个整数数组 nums，找到子数组（必须是连续的且至少包含一个元素）最大和。

#### 示例
* 输入 : [-2,1,-3,4,-1,2,1,-5,4]
* 输出 : 6
* 解释 : 子数组为 [4,-1,2,1] 和为 6
## 2. 解题思路
动态规划

设 dp[i] 表示已 nums[i] 结尾的子数组的最大和

那么 dp[i] = max(nums[i], dp[i-1] + nums[i])

如果 dp[i-1] 是负数，那么 dp[i] = nums[i]

如果 dp[i-1] 是正数或者 0 ，那么 dp[i] = dp[i-1] + nums[i]

初始化 dp[0] = nums[0]

有一个额外的变量记录最大值 max，每次得到 dp[i] 之后都与最大值 max 做比较 max=max(max,dp[i])

返回 max

## 3. 代码
```golang 代码
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
```
## 4. 复杂度分析
* 时间复杂度 : O(N) N 为 nums 数组长度，遍历一遍 nums
* 空间复杂度 : O(N) dp 的数组长度

---

# Review []()

---

# Tip


---

# Share
