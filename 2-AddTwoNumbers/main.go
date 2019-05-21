package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		l1 = &ListNode{Val: 2}
		l2 = &ListNode{Val: 5}
		lr = &ListNode{Val: 7}
	)
	makeList(l1, 4, 3)
	makeList(l2, 6, 4)
	//printList(l1)
	//printList(l2)
	makeList(lr, 0, 8)
	fmt.Println(compareOnList(addTwoNumbers(l1, l2), lr))
}

func makeList(h *ListNode, is ...int) {
	p := h
	for i := range is {
		p.Next = &ListNode{Val: is[i]}
		p = p.Next
	}
}

func printList(h *ListNode) {
	p := h
	for p.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	fmt.Println(p.Val, ". ")
}

func compareOnList(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lr := new(ListNode)
	p := lr

	var flag, sum int

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
		p.Next = new(ListNode)
		p = p.Next
	}

	return lr
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var lr *ListNode = new(ListNode)
	p := lr

	var sum, flag int
	for {
		sum = l1.Val + l2.Val + flag
		p.Val = sum % 10
		flag = sum / 10
		p.Next = new(ListNode)
		if l1.Next == nil || l2.Next == nil {
			if flag == 1 {
				p.Next.Val = 1
			} else {
				p.Next = nil
			}
			break
		}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1.Next == nil && l2.Next == nil {
		return lr
	}

	sum = 0
	p.Next = new(ListNode)
	p = p.Next
	if l1.Next != nil {
		l1 = l1.Next
		for {
			sum = l1.Val + flag
			p.Val = sum % 10
			flag = sum / 10
			if l1.Next == nil {
				if flag == 1 {
					p.Next = new(ListNode)
					p.Next.Val = 1
				}
				break
			}
			p.Next = new(ListNode)
			p = p.Next
			l1 = l1.Next
		}
	}
	if l2.Next != nil {
		l2 = l2.Next
		for {
			sum = l2.Val + flag
			p.Val = sum % 10
			flag = sum / 10
			if l2.Next == nil {
				if flag == 1 {
					p.Next = new(ListNode)
					p.Next.Val = 1
				}
				break
			}
			p.Next = new(ListNode)
			p = p.Next
			l2 = l2.Next
		}
	}

	return lr
}
*/
