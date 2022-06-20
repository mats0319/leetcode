package mario

import "sort"

type DinnerPlates struct {
	stackSlice   []*stack
	notFullStack []int
	capacity     int
}

type stack struct {
	data []int
	top  int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		stackSlice:   make([]*stack, 0),
		notFullStack: make([]int, 0),
		capacity:     capacity,
	}
}

func (dp *DinnerPlates) Push(val int) {
	if len(dp.notFullStack) < 1 { // no valid stack, need a new one
		newStackData := make([]int, dp.capacity)
		newStackData[0] = val

		dp.stackSlice = append(dp.stackSlice, &stack{
			data: newStackData,
			top:  1,
		})
		if dp.capacity > 1 {
			dp.notFullStack = append(dp.notFullStack, len(dp.stackSlice)-1)
		}
	} else {
		stackIndex := dp.notFullStack[0]

		stackIns := dp.stackSlice[stackIndex]

		stackIns.data[stackIns.top] = val
		stackIns.top++

		dp.stackSlice[stackIndex] = stackIns

		if stackIns.top >= dp.capacity {
			dp.notFullStack = dp.notFullStack[1:]
		}
	}
}

// Pop if all stack are empty, return -1
func (dp *DinnerPlates) Pop() int {
	var (
		stackIns *stack
		index    int
	)

	for index = len(dp.stackSlice) - 1; index >= 0; index-- {
		if dp.stackSlice[index].top > 0 { // not empty
			stackIns = dp.stackSlice[index]
			break
		}
	}

	if stackIns == nil {
		return -1
	}

	if stackIns.top >= dp.capacity {
		dp.notFullStack = append(dp.notFullStack, index)
		sort.Ints(dp.notFullStack)
	}

	stackIns.top--

	res := stackIns.data[stackIns.top]

	return res
}

// PopAtStack if target stack is empty, return -1
func (dp *DinnerPlates) PopAtStack(index int) int {
	if index >= len(dp.stackSlice) {
		return -1
	}

	stackIns := dp.stackSlice[index]

	if stackIns.top < 1 {
		return -1
	}

	if stackIns.top >= dp.capacity {
		dp.notFullStack = append(dp.notFullStack, index)
		sort.Ints(dp.notFullStack)
	}

	stackIns.top--

	res := stackIns.data[stackIns.top]

	return res
}
