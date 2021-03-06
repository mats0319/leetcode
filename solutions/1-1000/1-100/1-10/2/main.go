package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	p := pre

	addFlag := 0
	for l1 != nil || l2 != nil || addFlag == 1 {
		sum := getValue(l1) + getValue(l2) + addFlag

		addFlag = sum / 10
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next

		l1 = next(l1)
		l2 = next(l2)
	}

	return pre.Next
}

func next(l *ListNode) (res *ListNode) {
	if l != nil {
		res = l.Next
	}

	return
}

func getValue(l *ListNode) (res int) {
	if l != nil {
		res = l.Val
	}

	return
}
