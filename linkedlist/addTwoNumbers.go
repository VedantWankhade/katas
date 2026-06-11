package linkedlist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, n2 := l1, l2
	carry := 0

	out := &ListNode{
		Val:  n1.Val + n2.Val,
		Next: nil,
	}
	if n1.Val+n2.Val >= 10 {
		carry = 1
		out.Val %= 10
	}

	res := out

	for n1, n2 = n1.Next, n2.Next; n1 != nil && n2 != nil; n1, n2 = n1.Next, n2.Next {
		out.Next = &ListNode{
			Val:  n1.Val + n2.Val + carry,
			Next: nil,
		}
		if out.Next.Val >= 10 {
			carry = 1
			out.Next.Val %= 10
		} else {
			carry = 0
		}
		out = out.Next
	}

	if n1 == nil {
		for n := n2; n != nil; n = n.Next {
			out.Next = &ListNode{
				Val:  n.Val + carry,
				Next: nil,
			}
			if out.Next.Val >= 10 {
				carry = 1
				out.Next.Val %= 10
			} else {
				carry = 0
			}
			out = out.Next
		}
	}
	if n2 == nil {
		for n := n1; n != nil; n = n.Next {
			out.Next = &ListNode{
				Val:  n.Val + carry,
				Next: nil,
			}
			if out.Next.Val >= 10 {
				carry = 1
				out.Next.Val %= 10
			} else {
				carry = 0
			}
			out = out.Next
		}
	}

	if carry != 0 {
		out.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return res
}
