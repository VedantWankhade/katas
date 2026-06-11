package linkedlist_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/linkedlist"
	"github.com/vedantwankhade/katas/leetcode/top-interview-150/utils"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1       *linkedlist.ListNode
		l2       *linkedlist.ListNode
		expected *linkedlist.ListNode
	}{
		{
			l1:       utils.LinkedListFromArray([]int{2, 4, 3}),
			l2:       utils.LinkedListFromArray([]int{5, 6, 4}),
			expected: utils.LinkedListFromArray([]int{7, 0, 8}),
		},
		{
			l1:       utils.LinkedListFromArray([]int{0}),
			l2:       utils.LinkedListFromArray([]int{0}),
			expected: utils.LinkedListFromArray([]int{0}),
		},
		{
			l1:       utils.LinkedListFromArray([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       utils.LinkedListFromArray([]int{9, 9, 9, 9}),
			expected: utils.LinkedListFromArray([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, test := range tests {
		if actual := linkedlist.AddTwoNumbers(test.l1, test.l2); !utils.LinkedListEqual(test.expected, actual) {
			t.Errorf("\nL1: %v\tL2: %v\nExpected: %v\tActual: %v\n", utils.LinkedListString(test.l1), utils.LinkedListString(test.l2), utils.LinkedListString(test.expected), utils.LinkedListString(actual))
		}

	}
}
