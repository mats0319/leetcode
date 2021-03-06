package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p, q := head, head
	for p != nil && k > 0 {
		p = p.Next
		k--
	}
	if k > 0 {
		return nil
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}

	return q
}
