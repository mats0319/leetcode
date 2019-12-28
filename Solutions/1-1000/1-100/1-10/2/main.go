package mario

import "github.com/mats9693/leetcode/utils/compare"

func addTwoNumbers(l1 *cmp.ListNode, l2 *cmp.ListNode) *cmp.ListNode {
	var (
		lr = new(cmp.ListNode)
		p  = lr

		flag, sum int
	)

	for {
		sum = flag
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Val = sum % 10
		flag = sum / 10

		if l1 == nil && l2 == nil && flag == 0 {
			break
		}
		p.Next = new(cmp.ListNode)
		p = p.Next
	}

	return lr
}