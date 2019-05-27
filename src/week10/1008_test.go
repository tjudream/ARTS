package week10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func bstFromPreorderNoRecursion(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	for i,num := range preorder {
		if i == 0 {
			continue
		}
		node := &TreeNode{num, nil, nil}
		tmp := root
		for tmp != nil {
			if num > tmp.Val {
				if tmp.Right == nil {
					tmp.Right = node
					tmp = nil
				} else {
					tmp = tmp.Right
				}
			} else {
				if tmp.Left == nil {
					tmp.Left = node
					tmp = nil
				} else {
					tmp = tmp.Left
				}
			}
		}
	}
	return root
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var root *TreeNode
	for _,num := range preorder {
		root = insertNode(num, root)
	}
	return root
}
func insertNode(num int, root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{num, nil, nil}
	}
	if (num > root.Val) {
		root.Right = insertNode(num, root.Right)
	} else {
		root.Left = insertNode(num, root.Left)
	}
	return root
}

func preorderTreeRec(root *TreeNode, preorder []int) []int {
	if root != nil {
		preorder = append(preorder, root.Val)
	}
	if root.Left != nil {
		preorder = preorderTreeRec(root.Left, preorder)
	}
	if root.Right != nil {
		preorder = preorderTreeRec(root.Right, preorder)
	}
	return preorder
}
func preorderTree(root *TreeNode) []int {
	var preorder []int
	return preorderTreeRec(root, preorder)
}
func TestBstFromPreorder(t *testing.T) {
	preorder := []int{8,5,1,7,10,12}
	root := bstFromPreorder(preorder)
	assert.Equal(t, preorder, preorderTree(root))
}
func TestBstFromPreorderRescursion(t *testing.T) {
	preorder := []int{8,5,1,7,10,12}
	root := bstFromPreorderNoRecursion(preorder)
	assert.Equal(t, preorder, preorderTree(root))
}
