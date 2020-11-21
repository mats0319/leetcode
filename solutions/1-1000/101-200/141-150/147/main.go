package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{Next: &ListNode{Val: head.Val}}
	head = head.Next
	for head != nil {
		p := pre

		for p.Next != nil && p.Next.Val < head.Val {
			p = p.Next
		}

		next := p.Next
		p.Next = &ListNode{
			Val:  head.Val,
			Next: next,
		}

		head = head.Next
	}

	return pre.Next
}
