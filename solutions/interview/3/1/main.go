package mario

type TripleInOne struct {
	data   []int
	length []int
	size   int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		data:   make([]int, stackSize*3),
		length: make([]int, 3),
		size:   stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	length := this.length[stackNum]
	if length >= this.size {
		return
	}

	this.data[stackNum*this.size+length] = value
	this.length[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	length := this.length[stackNum]
	if length <= 0 {
		return -1
	}

	this.length[stackNum]--

	return this.data[stackNum*this.size+length-1]
}

func (this *TripleInOne) Peek(stackNum int) int {
	length := this.length[stackNum]
	if length <= 0 {
		return -1
	}

	return this.data[stackNum*this.size+length-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.length[stackNum] <= 0
}
