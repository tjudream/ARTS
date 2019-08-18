package week22

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

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
