package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	p := head
	count := 1
	for p.Next != nil {
		p = p.Next
		count++
	}
	p.Next = head
	k %= count
	for i := count*2 - k; i > 0; i-- {
		p = p.Next
	}

	t := p.Next
	p.Next = nil

	return t
}
