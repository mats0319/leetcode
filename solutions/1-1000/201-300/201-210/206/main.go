package mario

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
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

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return revert(nil, head, head.Next) // todo: tail recursion optimize
}

func revert(pre, cur, next *ListNode) *ListNode {
	cur.Next = pre

	if next == nil {
		return cur
	}

	return revert(cur, next, next.Next)
}

func reverseListIteration(head *ListNode) *ListNode {
	if head == nil { // head.next == nil is ok
		return head
	}

	var pre, cur, next *ListNode // now, pre is nil
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
