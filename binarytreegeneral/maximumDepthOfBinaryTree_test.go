package binarytreegeneral_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/binarytreegeneral"
)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	tests := []struct {
		root     *binarytreegeneral.TreeNode
		expected int
	}{
		{
			root: &binarytreegeneral.TreeNode{
				Val: 3,
				Left: &binarytreegeneral.TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &binarytreegeneral.TreeNode{
					Val: 20,
					Left: &binarytreegeneral.TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &binarytreegeneral.TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			expected: 3,
		},
		{
			root: &binarytreegeneral.TreeNode{
				Val:  2,
				Left: nil,
				Right: &binarytreegeneral.TreeNode{
					Val:   2,
					Right: nil,
					Left:  nil,
				},
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		if actual := binarytreegeneral.MaximumDepthOfBinaryTree(test.root); actual != test.expected {
			t.Errorf("\nExpeted: %v\tActual: %v", test.expected, actual)
		}
	}
}

/*
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:

Input: root = [1,null,2]
Output: 2

*/
