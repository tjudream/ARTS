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
    
# Share 22 MySQL有哪些“饮鸩止渴”提高性能的方法？ —— 极客时间 MySQL实战45讲
## 短连接风暴
MySQL 建立连接的成本是很高的。

短连接模型中，如果数据库处理的慢一些，连接数就会暴涨。

max_connections 参数，用来控制 MySQL 实例同时存在的连接数的上限，超过这个值，系统会拒绝连接，并返回
“Too many connections”错误。

遇到这种情况，如果只是单纯的调大 max_connections 的值的话，是有风险的，MySQL 很有可能会耗尽系统资源，如 CPU、内存等

那么是否还有其他办法呢? 这里还有两种方法，但都是有损的。

### 第一种方法：先处理掉那些占着连接但是不工作的线程
* 可以通过 kill connection 主动踢掉
* 设置 wait_timeout，一个线程空闲 wait_timeout 这么多秒之后，就会被 MySQL 直接断开

在 show processlist 的结果中，踢掉 sleep 的线程，可能是有损的

| | session A | session B | session C |
| --- | --- | --- | --- |
| T | begin; <br> insert into t values(1,1); | select * from t where id=1; | |
| T +30s | | | show processlist; |



