package binarytreegeneral

// Problem: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaximumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaximumDepthOfBinaryTree(root.Left), MaximumDepthOfBinaryTree(root.Right)) + 1
}
