package mario

type NestedIterator2 struct {
	stack []*NestedInteger
}

func Constructor2(nestedList []*NestedInteger) *NestedIterator2 {
	stack := make([]*NestedInteger, 0, len(nestedList))
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}

	return &NestedIterator2{
		stack: stack,
	}
}

// a HasNext() will be caller before Next()
func (this *NestedIterator2) Next() int {
	value := this.stack[len(this.stack)-1].GetInteger()
	this.stack = this.stack[:len(this.stack)-1]

	return value
}

func (this *NestedIterator2) HasNext() bool {
	if len(this.stack) < 1 {
		return false
	} else if this.stack[len(this.stack)-1].IsInteger() {
		return true
	}

	// unfold list
	value := this.stack[len(this.stack)-1].GetList()
	this.stack = this.stack[:len(this.stack)-1]

	for i := len(value) - 1; i >= 0; i-- {
		this.stack = append(this.stack, value[i])
	}

	return this.HasNext()
}
