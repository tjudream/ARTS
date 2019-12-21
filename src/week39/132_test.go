package week39

const INT_MAX = int(^uint(0) >> 1)

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i:=0; i<n+1;i++ {
		dp[i] = INT_MAX
	}
	dp[0] = -1
	for i:=0; i<n; i++ {
		//palindrome of length 1,3,5...
		for len:=0;i-len>=0 && i+len<n && s[i-len] == s[i+len]; len++ {
			if dp[i+len+1] > dp[i-len]+1 {
				dp[i+len+1] = dp[i-len]+1
			}
		}
		//palindrome of length 2,4,6...
		for len:=0;i-len>=0 && i+len+1<n && s[i-len] == s[i+len+1]; len++ {
			if dp[i+len+2] > dp[i-len]+1 {
				dp[i+len+2]=dp[i-len]+1
			}
		}
	}
	return dp[n]
}
