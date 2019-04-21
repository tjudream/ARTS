package week4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	} else {
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	}
}

func addToBST(root *TreeNode, node *TreeNode) {
	if node.Val <= root.Val {
		if root.Left == nil {
			root.Left = node
		} else {
			addToBST(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			addToBST(root.Right, node)
		}
	}
}
func generateBST(arr []int) *TreeNode {
	root := &TreeNode{
		arr[0],
		nil,
		nil,
	}
	for _,v := range arr[1:] {
		node := &TreeNode{
		v,
		nil,
		nil,
		}
		addToBST(root, node)
	}
	return root
}
func TestRangeSumBST(t *testing.T) {
	arr := []int{10,5,15,3,7,18}
	root := generateBST(arr)
	assert.Equal(t, 32, rangeSumBST(root, 7, 15))
	arr = []int{10,5,15,3,7,13,18,1,6}
	root = generateBST(arr)
	assert.Equal(t, 23, rangeSumBST(root, 6, 10))
}