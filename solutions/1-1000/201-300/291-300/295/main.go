package mario

type MedianFinder struct {
	pre   *listNode
	left  *listNode
	length int
}

type listNode struct {
	value int
	prev  *listNode
	next  *listNode
}

func Constructor() MedianFinder {
	return MedianFinder{
		pre:   &listNode{},
		left:  nil,
		length: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.pre.insertInOrder(num)

	this.length++

	if this.left == nil {
		this.left = this.pre.next
		return
	}

	if num <= this.left.value && this.length % 2 == 0 {
		this.left = this.left.prev
	} else if num > this.left.value && this.length % 2 == 1 {
		this.left = this.left.next
	}
}

func (this *MedianFinder) FindMedian() float64 {
	right := 0
	if this.length % 2 == 1 {
		right = this.left.value
	} else {
		right = this.left.next.value
	}

	return float64(this.left.value+right) / 2.0
}

func (ln *listNode) insertInOrder(value int) {
	pre := ln

	// same value will add at the start
	for pre.next != nil && pre.next.value < value {
		pre = pre.next
	}

	next := pre.next
	newNode := &listNode{
		value: value,
		prev:  pre,
		next:  next,
	}

	pre.next = newNode
	if next != nil {
		next.prev = newNode
	}
}
