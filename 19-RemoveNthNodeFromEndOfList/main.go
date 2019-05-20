package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		head   = &ListNode{Val: 1}
		result = &ListNode{Val: 1}
	)
	makeList(head, 2, 3, 4, 5)
	makeList(result, 2, 3, 5)
	//printList(head)
	//printList(result)
	fmt.Println(compareOnLinkedList(removeNthFromEnd(head, 2), result))
}

func makeList(h *ListNode, is ...int) {
	for i := range is {
		connectNode(h, is[i])
	}
}
func connectNode(h *ListNode, i int) {
	var p *ListNode
	for p = h; p.Next != nil; p = p.Next {
	}
	p.Next = &ListNode{Val: i}
}

func printList(h *ListNode) {
	var p *ListNode
	for p = h; p.Next != nil; p = p.Next {
		fmt.Printf("%d -> ", p.Val)
	}
	fmt.Printf("%d.\n", p.Val)
}

func compareOnLinkedList(a, b *ListNode) (equal bool) {
	for a != nil {
		if a.Val != b.Val {
			break
		}
		a = a.Next
		b = b.Next
	}

	if a == nil && b == nil {
		equal = true
	}

	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		step int
		temp *ListNode
	)

	for temp = head; temp != nil; temp = temp.Next {
		step++
	}

	step -= n + 1 // step = step + 1 - n - 2, step is steps need to move.
	if step == -1 {
		return head.Next
	}

	for temp = head; step > 0; step-- {
		temp = temp.Next // temp point to pre of target node finally.
	}
	temp.Next = temp.Next.Next // delete the end is also ok.

	return head
}