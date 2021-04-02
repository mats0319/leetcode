package mario

type NestedIterator struct {
	array     []int
	nextIndex int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		array:     unfoldNested(nestedList),
		nextIndex: 0,
	}
}

func unfoldNested(list []*NestedInteger) []int {
	if len(list) < 1 {
		return nil
	}

	res := make([]int, 0)

	for i := 0; i < len(list); i++ {
		node := list[i]

		if node.IsInteger() {
			res = append(res, node.GetInteger())
		} else {
			res = append(res, unfoldNested(node.GetList())...)
		}
	}

	return res
}

func (this *NestedIterator) Next() int {
	value := this.array[this.nextIndex]
	this.nextIndex++

	return value
}

func (this *NestedIterator) HasNext() bool {
	return this.nextIndex < len(this.array)
}
