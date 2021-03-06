package mario

import "sort"

type listNode struct {
	value int
	next  *listNode
}

// k <= nums.length
func findKthLargest(nums []int, k int) int {
	list := makeList(nums[:k])
	for i := k; i < len(nums); i++ {
		if nums[i] > list.value {
			list = updateList(list, nums[i])
		}
	}

	return list.value
}

// list.value is from small to big
// list returned not contains pre-node
func makeList(elems []int) *listNode {
	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	pre := &listNode{}
	q := pre
	for _, v := range elems {
		q.next = &listNode{value: v}
		q = q.next
	}

	return pre.next
}

// value > head.value
func updateList(head *listNode, value int) *listNode {
	pre := &listNode{next: head}
	p := head
	for p != nil && value > p.value {
		pre = pre.next
		p = p.next
	}

	pre.next = &listNode{value: value, next: p}

	return head.next
}
