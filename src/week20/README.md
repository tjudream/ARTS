# week20

---

# Algorithm [814. Binary Tree Pruning](https://leetcode.com/problems/binary-tree-pruning/)
## 1. 问题描述
修剪二叉树

给一个二叉树的根 root，且二叉树中节点的值只包含 0,1

要求剪掉所不包含值为 1 的节点的所有子树，并返回新树的树根 root

## 2. 解题思路
后序遍历：
1. 递归计算左子树中所有节点的和，如果等于 0 ，则剪掉
2. 递归计算右子树中所有节点的和，如果等于 0， 则剪掉

## 3. 代码
```go
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
```
## 4. 复杂度分析
* 时间复杂度: O(N), N 为树中节点个数，需要全遍历一遍
* 空间复杂度: O(logN) 需要树的高度的栈深度

---

# Review []()

---

# Tip

## 

---
    
# Share
## 

