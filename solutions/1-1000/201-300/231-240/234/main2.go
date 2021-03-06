package mario

func isPalindromeS2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	rl := revertList(slow.Next)
	p := head

	isPalindromeFlag := true
	for rl != nil {
		if rl.Val != p.Val {
			isPalindromeFlag = false
			break
		}

		rl = rl.Next
		p = p.Next
	}

	//slow.Next = revertList(rl)

	return isPalindromeFlag
}

func revertList(l *ListNode) *ListNode {
	var pre *ListNode
	for l != nil {
		next := l.Next

		l.Next = pre
		pre = l
		l = next
	}

	return pre
}
