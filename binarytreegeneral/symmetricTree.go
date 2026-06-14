package binarytreegeneral

func SymmetricTree(head *TreeNode) bool {
	if head == nil {
		return true
	}

	return symmetricSides(head.Left, head.Right)
}

func symmetricSides(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return symmetricSides(left.Left, right.Right) && symmetricSides(left.Right, right.Left)
}
