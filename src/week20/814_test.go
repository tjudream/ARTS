package week20

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func prunceTree(root *TreeNode) *TreeNode {
	sum := sumTree(root)
	if sum == 0 {
		return nil
	}
	return root
}

func sumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := sumTree(root.Left)
	rightSum := sumTree(root.Right)
	if leftSum == 0 {
		root.Left = nil
	}
	if rightSum == 0 {
		root.Right = nil
	}
	return root.Val + leftSum + rightSum
}

