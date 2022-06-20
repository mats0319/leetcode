package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{Next: head}
	for p := result; p.Next != nil && p.Next.Next != nil; p = p.Next.Next {
		q := p.Next

		p.Next = q.Next
		q.Next = p.Next.Next
		p.Next.Next = q
	}

	return result.Next
}
