package mario

type MyCircularQueue struct {
	queue    []int
	length   int
	capacity int
	head     int
	tail     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:    make([]int, k),
		length:   0,
		capacity: k,
		head:     0,
		tail:     -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.tail = (this.tail + 1) % this.capacity

	this.queue[this.tail] = value
	this.length++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % this.capacity
	this.length--

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.length == this.capacity
}
