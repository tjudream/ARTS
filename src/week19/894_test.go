package week19

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	res := []*TreeNode{}
	if N % 2 == 0 {
		return res
	}
	if N == 1 {
		res = append(res, &TreeNode{0,nil,nil})
		return res
	}
	N -= 1
	for i := 1; i < N; i += 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(N - i)
		for _,nl := range left {
			for _,nr := range right {
				cur := &TreeNode{0,nl,nr}
				res = append(res, cur)
			}
		}
	}
	return res
}



