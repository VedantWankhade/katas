package linkedlist

// Problem: https://leetcode.com/problems/linked-list-cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func LinkedListCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}
	if head.Next.Next == nil {
		return false
	}
	if head.Next.Next == head {
		return true
	}
	if head.Next.Next == head.Next {
		return true
	}

	slow, fast := head, head.Next.Next
	for fast != nil && slow != fast {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
	}
	if fast == nil {
		return false
	}
	return true
}

func LinkedListCycleHashmap(head *ListNode) bool {
	hash := make(map[*ListNode]struct{})

	for n := head; n != nil; n = n.Next {
		if _, ok := hash[n]; ok {
			return true
		}
		hash[n] = struct{}{}
	}
	return false
}
