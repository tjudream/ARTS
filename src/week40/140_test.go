package week40

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func wordBreak(s string, wordDict []string) []string {
	return dfs(s, wordDict, map[string][]string{})
}

func dfs(s string, wordDict []string, m map[string][]string) []string {
	if val, ok := m[s]; ok {
		return val
	}

	if len(s) == 0 {
		return []string{""}
	}

	var res []string
	for _, w := range wordDict {
		if len(w) <= len(s) && w == s[:len(w)] {
			for _, str := range dfs(s[len(w):], wordDict, m) {
				if len(str) == 0 {
					res = append(res, w)
				} else {
					res = append(res, w+" "+str)
				}
			}
		}
	}
	m[s] = res
	return res
}

func TestWordBreak(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	res := []string{"cat sand dog", "cats and dog"}
	assert.Equal(t, res, wordBreak(s, wordDict))
}
