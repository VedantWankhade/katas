package binarytreegeneral

func PathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, targetSum)
}

func pathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val-targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return pathSum(root.Left, targetSum-root.Val) || pathSum(root.Right, targetSum-root.Val)
}
