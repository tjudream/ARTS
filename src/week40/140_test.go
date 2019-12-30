package week40

import "testing"

func wordBreak(s string, wordDict []string) []string {
	var m map[string][]string
	return dfs(s, wordDict, m)
}

func dfs(s string, wordDict []string, m map[string][]string) []string {

	return nil
}

func TestWordBreak(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	res := []string{"cats and dog", "cat sand dog"}
	assert.Equals(t, res, wordBreak(s, wordDict))
}
