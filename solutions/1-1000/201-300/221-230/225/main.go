package mario

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	length := len(this.queue)

	this.queue = append(this.queue, x)
	for i := 0; i < length; i++ {
		this.queue = append(this.queue, this.Pop())
	}

	return
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	value := this.queue[0]
	this.queue = this.queue[1:]
	return value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
