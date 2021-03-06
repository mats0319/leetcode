package mario

type MyQueue struct {
	receive *stack
	output  *stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		receive: newStack(),
		output:  newStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.receive.push(x)
}

// valid caller, unnecessary error handler
/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.output.isEmpty() {
		// 题目已声明，可以不考虑错误处理，即调用该函数时，队列内必不为空
		for !this.receive.isEmpty() {
			this.output.push(this.receive.pop())
		}
	}

	return this.output.pop()
}

// valid caller, unnecessary error handler
/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.output.isEmpty() {
		// 题目已声明，可以不考虑错误处理，即调用该函数时，队列内必不为空
		for !this.receive.isEmpty() {
			this.output.push(this.receive.pop())
		}
	}

	return this.output.array[len(this.output.array)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.receive.isEmpty() && this.output.isEmpty()
}

type stack struct {
	array []int
}

func newStack() *stack {
	return &stack{
		array: make([]int, 0),
	}
}

func (s *stack) push(elem int) {
	s.array = append(s.array, elem)
}

func (s *stack) pop() int {
	value := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]

	return value
}

func (s *stack) isEmpty() bool {
	return len(s.array) < 1
}
