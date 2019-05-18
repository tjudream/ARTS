#week8

---

## Algorithm [701. Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/)
### 1. 问题描述
将一个元素插入到二叉搜索树中。
给定二叉搜索树的根 root，且新元素在树中不存在。
### 2. 解题思路
二叉索索树定义：
* 左子树中所有节点的值都小于此节点
* 右子树中所有节点的值都大于此节点
* 左右子树分别都是二叉搜索树

将新元素与当前节点的值对比，如果大于当前节点则插入到其右子树中，如果小于当前节点则插入其左子树中。
递归调用直到叶子节点为止。

### 3. 代码
#### 3.1 递归
```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
```
#### 3.2 非递归
```go
func insertIntoBSTWithoutRecursion(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	node := root
	for node != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{val, nil, nil}
				break
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{val, nil, nil}
				break
			} else {
				node = node.Right
			}
		}
	}
	return root
}
```
### 4. 复杂度分析
* 时间复杂度 O(H), 其中 H = logN 为树的高度。 最坏情况为 O(N) 二叉树退化为链表的情况。
* 空间复杂度:
    * 递归： O(H) 栈的深度
    * 非递归： O(1) 只需要一个节点的空间存储临时节点。

---

## Review []()

---

## Tip

### 

---
    
## Share
### 

