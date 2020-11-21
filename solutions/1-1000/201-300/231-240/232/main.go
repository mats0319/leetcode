package mario

type MyQueue struct {
	stackOne []int
	stackTwo []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackOne: make([]int, 0),
		stackTwo: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackOne = append(this.stackOne, x)
}

// valid caller, unnecessary error handler
/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stackTwo) > 0 {
		v := this.stackTwo[len(this.stackTwo)-1]
		this.stackTwo = this.stackTwo[:len(this.stackTwo)-1]

		return v
	}

	for i := len(this.stackOne)-1; i > 0; i-- {
		this.stackTwo = append(this.stackTwo, this.stackOne[i])
	}

	v := this.stackOne[0]
	this.stackOne = make([]int, 0)

	return v
}

// valid caller, unnecessary error handler
/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackTwo) > 0 {
		return this.stackTwo[len(this.stackTwo)-1]
	}

	return this.stackOne[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stackOne)+len(this.stackTwo) < 1
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
