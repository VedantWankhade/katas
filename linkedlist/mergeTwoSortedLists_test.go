package linkedlist_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/linkedlist"
	"github.com/vedantwankhade/katas/leetcode/top-interview-150/utils"
)

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		list1    *linkedlist.ListNode
		list2    *linkedlist.ListNode
		expected *linkedlist.ListNode
	}{
		{
			list1: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val: 2,
					Next: &linkedlist.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			list2: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val: 3,
					Next: &linkedlist.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val: 1,
					Next: &linkedlist.ListNode{
						Val: 2,
						Next: &linkedlist.ListNode{
							Val: 3,
							Next: &linkedlist.ListNode{
								Val: 4,
								Next: &linkedlist.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			list1: nil,
			list2: &linkedlist.ListNode{
				Val:  0,
				Next: nil,
			},
			expected: &linkedlist.ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}
	for _, test := range tests {
		if actual := linkedlist.MergeTwoSortedLists(test.list1, test.list2); !utils.LinkedListEqual(actual, test.expected) {
			t.Errorf("\nExpected: %v\tActual: %v", utils.LinkedListString(test.expected), utils.LinkedListString(actual))
		}
	}
}
