package mario

import (
	"github.com/mats9693/utils/compare"
)

func reverseList(head *cmp.ListNode) *cmp.ListNode {
	stack := make([]int, 0)

	p := head
	for ; p != nil; p = p.Next {
		stack = append(stack, p.Val)
	}

	p = head
	for i := len(stack) - 1; i >= 0; i-- {
		p.Val = stack[i]
		p = p.Next
	}

	return head
}

func reverseListRecursive(head *cmp.ListNode) *cmp.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return revert(nil, head, head.Next)
}

func revert(pre, cur, next *cmp.ListNode) *cmp.ListNode {
	cur.Next = pre

	if next == nil {
		return cur
	}

	return revert(cur, next, next.Next)
}

func reverseListIteration(head *cmp.ListNode) *cmp.ListNode {
	if head == nil { // head.next == nil is ok
		return head
	}

	var pre, cur, next *cmp.ListNode // now, pre is nil
	for cur, next = head, head.Next; ; next = next.Next {
		cur.Next = pre

		if next == nil { // if judge and break is at the end of loop, function need to return 'pre'.
			break
		}

		pre = cur
		cur = next
	}

	return cur
}
