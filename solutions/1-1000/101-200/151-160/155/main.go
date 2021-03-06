package mario

import "math"

type MinStack struct {
	stack []int
	small []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 1),
		small: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.small = append(this.small, min(this.small[len(this.small)-1], x))
}

func (this *MinStack) Pop() {
	length := len(this.stack)
	this.stack = this.stack[:length-1]
	this.small = this.small[:length-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.small[len(this.small)-1]
}

func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
