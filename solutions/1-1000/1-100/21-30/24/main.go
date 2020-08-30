package mario

import "github.com/mats9693/utils/compare"

func swapPairs(head *cmp.ListNode) *cmp.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := &cmp.ListNode{Next: head}
	for p := result; p.Next != nil && p.Next.Next != nil; p = p.Next.Next {
		q := p.Next

		p.Next = q.Next
		q.Next = p.Next.Next
		p.Next.Next = q
	}

	return result.Next
}
