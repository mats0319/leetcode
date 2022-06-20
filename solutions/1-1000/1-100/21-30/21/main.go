package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		p      = &ListNode{}
		result = p
	)

	for {
		if l1 == nil {
			p.Next = l2
			break
		} else if l2 == nil {
			p.Next = l1
			break
		}

		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}

		p = p.Next
	}

	return result.Next
}
