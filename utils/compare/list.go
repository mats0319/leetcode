package compare

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func compareOnList(a, b *ListNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	for a.Next != nil && b.Next != nil {
		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return a.Val == b.Val && a.Next == nil && b.Next == nil
}

func makeList(vals ...int) *ListNode {
	result := &ListNode{}
	p := result

	for i := range vals {
		p.Next = &ListNode{Val: vals[i]}
		p = p.Next
	}

	return result.Next
}

func printList(l *ListNode) {
	p := l

	for l.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	fmt.Printf("%d. ", p.Val)

	return
}
