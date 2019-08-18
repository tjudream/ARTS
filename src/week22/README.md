# week22

---

# Algorithm [1161. Maximum Level Sum of a Binary Tree](https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/)
## 1. 问题描述
找到二叉树中层中节点和最大的那一层

给定一个二叉树的根 root，root 的层是 1，root 的左右儿子的层是 2，以此类推。

将节点按照层进行分类，计算每一层所有节点的和，返回和最大的那个层的层号，如果有和相同的层，则返回最小的那层的层号。

示例 1：
![capture](capture.jpeg)
* 输入：[1,7,0,7,-8,null,null]
* 输出：2
* 解释：
    * 第 1 层 sum = 1
    * 第 2 层 sum = 7 + 0 = 7
    * 第 3 层 sum = 7 + (-8) = -1
    * 所以和最大的层的层号为 2

## 2. 解题思路
广度优先搜索，计算每一层的所有节点的和。

## 3. 代码
```go
type Element struct {
	node *TreeNode
	Level int
}
func maxLevelSum(root *TreeNode) int {
    if root == nil {
		return 0
	}
	maxSum := 0
	maxLevel := 1
	curSum := 0
	curLevel := 1
	elements := []*Element{&Element{root, 1}}

	for len(elements) > 0 {
		ele := elements[0]
		elements = elements[1:]
		node := ele.node
		level := ele.Level
		if node.Left != nil {
			elements = append(elements, &Element{node.Left, level + 1})
		}
		if node.Right != nil {
			elements = append(elements, &Element{node.Right, level + 1})
		}

		if level == curLevel {
			curSum += node.Val
		} else {
			if curSum > maxSum {
				maxSum = curSum
				maxLevel = curLevel
			}
			curSum = node.Val
			curLevel = level
		}
	}

	return maxLevel
}
```
## 4. 复杂度分析
* 时间复杂度： O(N) N 为树的节点个数，只需遍历一遍树的所有节点
* 空间复杂度： O(2<sup>H</sup>) H 为树高，root 的高为 1，因为数组 elements 最多需要存储一层的所有节点。
一个满二叉树的最后一层的节点数最多，为 2<sup>H</sup> 个。

---

# Review []()

---

# Tip

## 

---
    
# Share 24 MySQL是怎么保证主备一致的？ —— 极客时间 MySQL实战45讲
## 

