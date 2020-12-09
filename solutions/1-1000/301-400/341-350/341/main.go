package mario


// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
    panic("implement me")
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
    panic("implement me")
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
    panic("implement me")
}


type NestedIterator struct {
    array []int
    nextIndex int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
        array: unfoldNested(nestedList),
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
