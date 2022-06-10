package mario

type CountIntervals struct {
	pre *listNode
}

type listNode struct {
	left  int
	right int

	next *listNode
}

func Constructor() CountIntervals {
	return CountIntervals{
		pre: &listNode{},
	}
}

func (c *CountIntervals) Add(left int, right int) {
	pre := c.pre
	p := pre.next
	for p != nil && left > p.right {
		pre = p
		p = p.next
	}

	// 'left' is bigger than last interval, append new node at tail
	if p == nil {
		pre.next = &listNode{
			left:  left,
			right: right,
		}

		return
	}

	newNode := &listNode{
		left: small(left, p.left),
	}

	pre.next = newNode

	for p != nil && right > p.right {
		p = p.next
	}

	// 'right' is bigger than last interval, or smaller than p.left - 1
	if p == nil || right < p.left-1 {
		newNode.right = right
		newNode.next = p
	} else { // p.left - 1 <= 'right' <= p.right + 1
		newNode.right = big(right, p.right)
		newNode.next = p.next
	}

	return
}

func (c *CountIntervals) Count() int {
	res := 0
	for p := c.pre.next; p != nil; p = p.next {
		res += p.right - p.left + 1
	}

	return res
}

func small(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}

func big(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}

	return
}
