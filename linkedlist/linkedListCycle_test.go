package linkedlist_test

import (
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/linkedlist"
)

func TestLinkedListCycle(t *testing.T) {
	list1 := []int{3, 2, 0, -4}
	head1 := &linkedlist.ListNode{
		Val: 3,
	}
	head1.Next = &linkedlist.ListNode{
		Val: 2,
		Next: &linkedlist.ListNode{
			Val: 0,
			Next: &linkedlist.ListNode{
				Val:  -4,
				Next: head1,
			},
		},
	}

	list2 := []int{1, 2}
	head2 := &linkedlist.ListNode{
		Val: 1,
	}
	head2.Next = &linkedlist.ListNode{
		Val:  2,
		Next: head2,
	}

	list3 := []int{1}
	head3 := &linkedlist.ListNode{
		Val:  1,
		Next: nil,
	}
	tests := []struct {
		list     []int
		head     *linkedlist.ListNode
		expected bool
	}{
		{
			list:     list1,
			head:     head1,
			expected: true,
		},
		{
			list:     list2,
			head:     head2,
			expected: true,
		},
		{
			list:     list3,
			head:     head3,
			expected: false,
		},
	}

	for _, test := range tests {
		if actual := linkedlist.LinkedListCycle(test.head); actual != test.expected {
			t.Errorf("\nList: %v\nExpected: %v\tActual: %v\n", test.list, test.expected, actual)
		}
	}
}

func TestLinkedListCycleHashmap(t *testing.T) {
	list1 := []int{3, 2, 0, -4}
	head1 := &linkedlist.ListNode{
		Val: 3,
	}
	head1.Next = &linkedlist.ListNode{
		Val: 2,
		Next: &linkedlist.ListNode{
			Val: 0,
			Next: &linkedlist.ListNode{
				Val:  -4,
				Next: head1,
			},
		},
	}

	list2 := []int{1, 2}
	head2 := &linkedlist.ListNode{
		Val: 1,
	}
	head2.Next = &linkedlist.ListNode{
		Val:  2,
		Next: head2,
	}

	list3 := []int{1}
	head3 := &linkedlist.ListNode{
		Val:  1,
		Next: nil,
	}
	tests := []struct {
		list     []int
		head     *linkedlist.ListNode
		expected bool
	}{
		{
			list:     list1,
			head:     head1,
			expected: true,
		},
		{
			list:     list2,
			head:     head2,
			expected: true,
		},
		{
			list:     list3,
			head:     head3,
			expected: false,
		},
	}

	for _, test := range tests {
		if actual := linkedlist.LinkedListCycleHashmap(test.head); actual != test.expected {
			t.Errorf("\nList: %v\nExpected: %v\tActual: %v\n", test.list, test.expected, actual)
		}
	}
}
