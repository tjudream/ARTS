package week21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func partitionLabels(S string) []int {
	res := []int{}
	l := len(S)
	last := make([]int,26)
	for i := 0; i < l; i++ {
		last[S[i] - 'a'] = i
	}
	start,end := 0,0
	for i := 0; i < l; i++ {
		if last[S[i] - 'a'] > end {
			end = last[S[i] - 'a']
		}
		if end == i {
			r := end - start + 1
			start = end + 1
			res = append(res, r)
		}
	}
	return res
}

func TestPartitionLabels(t *testing.T)  {
	S := "ababcbacadefegdehijhklij"
	assert.Equal(t, []int{9,7,8}, partitionLabels(S))
}

