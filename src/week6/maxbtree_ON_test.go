package week6_test

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)
func New() *Stack {
	return &Stack{nil, 0}
}
func (this *Stack) Len() int {
	return this.length
}
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var stack *Stack
	var root *TreeNode
	for _, num := range nums {
		cur := &TreeNode{num, nil, nil}
		if root.Val < num {
			root = cur
		}
		for stack.length > 0 {
			sNode,_ := stack.Peek().(*TreeNode)
			if sNode.Val < num {
				cur.Left = sNode
			} else {
				sNode.Right = cur
			}
		}
	}
	return root
}
