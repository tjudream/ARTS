package week34

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func merge(root *TreeNode, t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if (root == nil) {
        root = new(TreeNode)
    }
    if (t1 != nil && t2 != nil) {
        root.Val = t1.Val + t2.Val
        root.Left = merge(root.Left, t1.Left, t2.Left)
        root.Right = merge(root.Right, t1.Right, t2.Right)
    } else if (t1 == nil && t2 != nil) {
        root.Val = t2.Val
        root.Left = merge(root.Left, nil, t2.Left)
        root.Right = merge(root.Right, nil, t2.Right)
    } else if (t1 != nil && t2 == nil) {
        root.Val = t1.Val
        root.Left = merge(root.Left, t1.Left, nil)
        root.Right = merge(root.Right, t1.Right, nil)
    } else {
        return nil
    }
    return root
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    return merge(nil, t1, t2)
}
