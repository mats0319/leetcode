package mario

import "sort"

type listNode struct {
	value     int
	frequency int
	next      *listNode
}

func topKFrequent(nums []int, k int) []int {
	var array []*listNode
	{
		m := make(map[int]int) // map, element - frequency
		for i := range nums {
			m[nums[i]]++
		}

		array = make([]*listNode, 0, len(m))
		for key, v := range m {
			array = append(array, &listNode{value: key, frequency: v})
		}
	}

	list := makeList(array[:k])
	for i := k; i < len(array); i++ {
		if array[i].frequency > list.frequency {
			list = updateList(list, array[i])
		}
	}

	res := make([]int, 0, k)
	for list != nil {
		res = append(res, list.value)
		list = list.next
	}

	return res
}

// sort by frequency, from small to big
func makeList(elems []*listNode) *listNode {
	sort.Slice(elems, func(i, j int) bool {
		return elems[i].frequency < elems[j].frequency
	})

	pre := &listNode{}
	q := pre
	for _, item := range elems {
		q.next = &listNode{value: item.value, frequency: item.frequency}
		q = q.next
	}

	return pre.next
}

// item.frequency > head.frequency
func updateList(head *listNode, item *listNode) *listNode {
	pre := &listNode{next: head}
	p := head
	for p != nil && item.frequency > p.frequency {
		pre = pre.next
		p = p.next
	}

	pre.next = item
	item.next = p

	return head.next
}
