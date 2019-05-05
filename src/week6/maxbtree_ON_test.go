package week6_test

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var vecs []*TreeNode
	for _, num := range nums {
		cur := &TreeNode{num, nil, nil}
		for len(vecs) > 0 && vecs[len(vecs) - 1].Val < num {
			cur.Left = vecs[len(vecs) - 1]
			vecs = vecs[:len(vecs) - 1]
		}
		if len(vecs) > 0 {
			vecs[len(vecs) -1].Right = cur
		}
		vecs = append(vecs, cur)
	}
	return vecs[0]
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3,2,1,6,0,5}
	fmt.Println(constructMaximumBinaryTree(nums))
}