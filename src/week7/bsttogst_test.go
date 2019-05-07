package week7

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"testing"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func bstToGstWithVal(root *TreeNode, val int) int {
	if root == nil {
		return val
	}
	if root.Right != nil {
		val = bstToGstWithVal(root.Right, val)
	}
	val += root.Val
	root.Val = val
	if root.Left != nil {
		val = bstToGstWithVal(root.Left, root.Val)
	}
	return val
}

func bstToGst(root *TreeNode) *TreeNode {
	bstToGstWithVal(root, 0)
	return root;
}

func insertTreeNode(data int, root *TreeNode) *TreeNode {
	if root == nil {
		root = &TreeNode{data, nil, nil}
		return root
	}
	if data < root.Val {
		root.Left = insertTreeNode(data, root.Left)
	} else {
		root.Right = insertTreeNode(data, root.Right)
	}
	return root
}

func constructBst(nums []int) *TreeNode {
	var root *TreeNode
	for _,num := range nums {
		root = insertTreeNode(num, root)
	}
	return root
}

func bfs(root *TreeNode) {
	q := queue.New()
	q.Enqueue(root)
	for q.Len() > 0 {
		node := q.Dequeue().(*TreeNode)
		fmt.Printf("%d,", node.Val)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}
}

func bstToGstTest(nums []int) {
	bst := constructBst(nums)
	bfs(bst)
	fmt.Println("")
	gst := bstToGst(bst)
	bfs(gst)
	fmt.Println("")
}
func TestBstToGst(t *testing.T) {
	nums1 := []int{4,1,6,0,2,5,7,3,8}
	nums2 := []int{0,1}
	bstToGstTest(nums1)
	bstToGstTest(nums2)
}
