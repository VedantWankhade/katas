package binarytreegeneral

func CountCompleteTreeNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + CountCompleteTreeNodes(root.Left) + CountCompleteTreeNodes(root.Right)
}
