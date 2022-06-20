package mario

type MyCircularDeque struct {
	head     *listNode
	tail     *listNode
	length   int
	capacity int
}

type listNode struct {
	value int
	prev  *listNode
	next  *listNode
}

// Constructor Initialize your data structure here. Set the size of the deque to be k.
func Constructor(k int) MyCircularDeque {
	headNode := &listNode{}
	tailNode := &listNode{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return MyCircularDeque{
		head:     headNode,
		tail:     tailNode,
		capacity: k,
	}
}

// InsertFront Adds an item at the front of Deque. Return true if the operation is successful.
func (mcd *MyCircularDeque) InsertFront(value int) bool {
	if mcd.IsFull() {
		return false
	}

	next := mcd.head.next

	node := &listNode{
		value: value,
	}

	mcd.head.next = node
	node.next = next

	next.prev = node
	node.prev = mcd.head

	mcd.length++

	return true
}

// InsertLast Adds an item at the rear of Deque. Return true if the operation is successful.
func (mcd *MyCircularDeque) InsertLast(value int) bool {
	if mcd.IsFull() {
		return false
	}

	pre := mcd.tail.prev

	node := &listNode{
		value: value,
	}

	pre.next = node
	node.next = mcd.tail

	mcd.tail.prev = node
	node.prev = pre

	mcd.length++

	return true
}

// DeleteFront Deletes an item from the front of Deque. Return true if the operation is successful.
func (mcd *MyCircularDeque) DeleteFront() bool {
	if mcd.IsEmpty() {
		return false
	}

	del := mcd.head.next

	mcd.head.next = del.next
	del.next.prev = mcd.head

	mcd.length--

	return true
}

// DeleteLast Deletes an item from the rear of Deque. Return true if the operation is successful.
func (mcd *MyCircularDeque) DeleteLast() bool {
	if mcd.IsEmpty() {
		return false
	}

	del := mcd.tail.prev

	del.prev.next = mcd.tail
	mcd.tail.prev = del.prev

	mcd.length--

	return true
}

// GetFront Get the front item from the deque.
func (mcd *MyCircularDeque) GetFront() int {
	if mcd.IsEmpty() {
		return -1
	}

	return mcd.head.next.value
}

// GetRear Get the last item from the deque.
func (mcd *MyCircularDeque) GetRear() int {
	if mcd.IsEmpty() {
		return -1
	}

	return mcd.tail.prev.value
}

// IsEmpty Checks whether the circular deque is empty or not.
func (mcd *MyCircularDeque) IsEmpty() bool {
	return mcd.head.next == mcd.tail
}

// IsFull Checks whether the circular deque is full or not.
func (mcd *MyCircularDeque) IsFull() bool {
	return mcd.length >= mcd.capacity
}
