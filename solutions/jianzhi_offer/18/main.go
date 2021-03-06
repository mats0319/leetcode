package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	pre := &ListNode{Next: head}

	for p := pre; p.Next != nil; p = p.Next {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
	}

	return pre.Next
}
