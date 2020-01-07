package mario

import "github.com/mats9693/utils/compare"

func mergeTwoLists(l1 *cmp.ListNode, l2 *cmp.ListNode) *cmp.ListNode {
	var (
		p      = &cmp.ListNode{}
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
