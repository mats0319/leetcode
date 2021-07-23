package mario

import "math"

type KthLargest struct {
	data *listNode // small -> big, same value update at first
}

type listNode struct {
	value int
	next  *listNode
}

func Constructor(k int, nums []int) KthLargest {
	pre := &listNode{value: math.MinInt64}
	p := pre
	for i := 0; i < k; i++ {
		p.next = &listNode{value: math.MinInt64}
		p = p.next
	}

	for i := range nums {
		pre.update(nums[i])
	}

	return KthLargest{data: pre.next}
}

func (this *KthLargest) Add(val int) int {
	this.data.update(val)

	return this.data.value
}

func (ln *listNode) update(value int) {
	if ln.value >= value {
		return
	}

	pre := ln
	for pre.next != nil && pre.next.value < value {
		pre = pre.next
	}

	pre.next = &listNode{
		value: value,
		next: pre.next,
	}

	*ln = *(ln.next)
}
