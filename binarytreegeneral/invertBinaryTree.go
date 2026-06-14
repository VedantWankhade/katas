package binarytreegeneral

func InvertBinaryTree(head *TreeNode) *TreeNode {
	if head == nil {
		return nil
	}

	head.Left = InvertBinaryTree(head.Left)
	head.Right = InvertBinaryTree(head.Right)

	head.Left, head.Right = head.Right, head.Left

	return head
}
