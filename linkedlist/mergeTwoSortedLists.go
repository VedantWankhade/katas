package linkedlist

func MergeTwoSortedLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	n1, n2 := list1, list2
	out := n1

	if n2.Val < n1.Val {
		out = n2
		n2 = n2.Next
	} else {
		n1 = n1.Next
	}

	res := out

	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			out.Next = n1
			n1 = n1.Next
		} else {
			out.Next = n2
			n2 = n2.Next
		}
		out = out.Next
	}

	if n1 == nil {
		for n := n2; n != nil; n = n.Next {
			out.Next = n
			out = out.Next
		}
	}

	if n2 == nil {
		for n := n1; n != nil; n = n.Next {
			out.Next = n
			out = out.Next
		}
	}

	return res
}
