package week6

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findMaxIndex(nums []int) int {
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 1 {
		return &TreeNode{nums[0],nil, nil}
	}
	maxIndex := findMaxIndex(nums)
	val := nums[maxIndex]
	var left, right *TreeNode
	if maxIndex > 0 {
		left = constructMaximumBinaryTree(nums[0:maxIndex])
	}
	if maxIndex + 1 < len(nums) {
		right = constructMaximumBinaryTree(nums[maxIndex + 1: len(nums)])
	}
	return &TreeNode{val, left, right}
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3,2,1,6,0,5}
	fmt.Println(constructMaximumBinaryTree(nums))
}
