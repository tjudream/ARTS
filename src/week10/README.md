#week10

---

## Algorithm [1008. Construct Binary Search Tree from Preorder Traversal](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)
### 1. 问题描述
用先序遍历数组构造二叉搜索树。

给定一个二叉搜索树的先序遍历数组，构造二叉搜索树。
### 2. 解题思路
先序遍历，即先遍历根，然后遍历左子树，最后遍历右子树。所以数组的第一个元素即是根。

构造出树的根之后，剩下的元素可以采用递归插入，或者非递归插入。
### 3. 代码
* 递归
```go
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
```
* 非递归
```go
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
```
### 4. 复杂度分析
* 时间复杂度: O(nlogn), n为数组中元素个数，遍历数组 n， 插入logn
* 空间复杂度: O(n), n为数组中元素个数

---

## Review []()

---

## Tip

### 

---
    
## Share
### 

