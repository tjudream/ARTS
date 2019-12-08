package week37

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func generate(s int, e int) []*TreeNode {
	var res []*TreeNode
	if s > e {
		res = append(res, nil)
		return res
	}
	if s == e {
		root := &TreeNode{s, nil, nil}
		res = append(res, root)
		return res
	}
	for i := s; i <= e; i++ {
		leftTrees := generate(s, i - 1)
		rightTrees := generate(i + 1, e)
		for l := 0; l < len(leftTrees); l++ {
			for r := 0; r < len(rightTrees); r++ {
				node := &TreeNode{i, leftTrees[l], rightTrees[r]}
				res = append(res, node)
			}
		}
	}
	return res
}

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}
	return generate(1, n)
}



