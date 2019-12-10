package mario

import "github.com/mats9693/leetcode/utils/compare"

func removeNthFromEnd(head *compare.ListNode, n int) *compare.ListNode {
	result := &compare.ListNode{Next: head}

	p, q := result, result
	for i := 0; i <= n; i++ {
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next

	return result.Next
}
