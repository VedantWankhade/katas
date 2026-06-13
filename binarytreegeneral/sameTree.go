package binarytreegeneral

func SameTree(p, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	leftSame := SameTree(p.Left, q.Left)
	rightSame := SameTree(p.Right, q.Right)

	return leftSame && rightSame
}
