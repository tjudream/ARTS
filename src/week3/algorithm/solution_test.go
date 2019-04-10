package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeOuterParentheses(S string) string {
	strArr := []rune(S)
	var sta,ret []rune
	for i := 0; i < len(S); i++ {
		switch strArr[i] {
		case '(':
			if len(sta) > 0 {
				ret = append(ret, strArr[i])
			}
			sta = append(sta, strArr[i])
		case ')':
			sta = sta[:len(sta) - 1]
			if len(sta) > 0 {
				ret = append(ret, strArr[i])
			}
		}
	}
	return string(ret)
}
func TestRemoveOuterParentheses(t *testing.T) {
	assert.Equal(t, "()()()", removeOuterParentheses("(()())(())"))
	assert.Equal(t, "()()()()(())", removeOuterParentheses("(()())(())(()(()))"))
	assert.Equal(t, "", removeOuterParentheses("()()"))
}
