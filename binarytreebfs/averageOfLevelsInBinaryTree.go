package binarytreebfs

import "github.com/vedantwankhade/katas/leetcode/top-interview-150/binarytreegeneral"

func AverageOfLevelsInBinaryTree(root *binarytreegeneral.TreeNode) []float64 {
	q := []*binarytreegeneral.TreeNode{root}
	out := []float64{}

	for len(q) > 0 {
		nums := len(q)
		sum := 0

		for range nums {
			n := q[0]
			q = q[1:]

			sum += n.Val

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		out = append(out, float64(sum)/float64(nums))
	}

	return out
}
