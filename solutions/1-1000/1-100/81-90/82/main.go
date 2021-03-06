package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{Next: head}
	p := head
	currVal := head.Val
	preNode := pre
	for p != nil && p.Next != nil {
		currVal = p.Val
		p = p.Next

		if p.Val == currVal {
			for p != nil && p.Val == currVal {
				p = p.Next
			}

			preNode.Next = p
		} else {
			preNode = preNode.Next
		}
	}

	return pre.Next
}
