package week26

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	if s == "" || p == "" {
		return false
	}
	dp := [][]bool{}

	dp[0][0] = true
	for 
}
