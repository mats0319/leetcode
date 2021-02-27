package mario

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	Iterator *Iterator
	PeekList    []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		Iterator: iter,
		PeekList: make([]int, 0),
	}
}

func (this *PeekingIterator) hasNext() bool {
	return len(this.PeekList) > 0 || this.Iterator.hasNext()
}

func (this *PeekingIterator) next() int {
	if len(this.PeekList) > 0 {
		v := this.PeekList[0]
		this.PeekList = this.PeekList[1:]
		return v
	} else if this.Iterator.hasNext() {
		return this.Iterator.next()
	}

	return -1
}

func (this *PeekingIterator) peek() int {
	if len(this.PeekList) > 0 {
		return this.PeekList[0]
	} else if this.Iterator.hasNext() {
		this.PeekList = append(this.PeekList, this.Iterator.next())
		return this.PeekList[0]
	}

	return -1
}
