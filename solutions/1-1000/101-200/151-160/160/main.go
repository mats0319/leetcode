package mario

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	stackA := make([]*ListNode, 0)
	for p := headA; p != nil; p = p.Next {
		stackA = append(stackA, p)
	}

	stackB := make([]*ListNode, 0)
	for p := headB; p != nil; p = p.Next {
		stackB = append(stackB, p)
	}

	ia, ib := len(stackA)-1, len(stackB)-1
	var firstIntersection *ListNode
	for ia >= 0 && ib >= 0 {
		if stackA[ia] == stackB[ib] {
			firstIntersection = stackA[ia]
		}

		ia--
		ib--
	}

	return firstIntersection
}
