package mario

// RangeModule
// 1 <= left < right < 10^9
type RangeModule struct {
	pre *listNode
}

type listNode struct {
	start int
	end   int
	next  *listNode
}

func Constructor() RangeModule {
	return RangeModule{
		pre: &listNode{},
	}
}

// AddRange 为了统一处理区间左、右端点是否在同一个区间段（包括一个跟踪区间及其左侧的未跟踪区间）内，新增区间总是引入一个新的节点
func (r *RangeModule) AddRange(left int, right int) {
	if r.pre.next == nil {
		r.pre.next = &listNode{
			start: left,
			end:   right,
		}

		return
	}

	pre := r.pre
	p := pre.next

	// find 'left' position
	for p != nil && left > p.end {
		pre = p
		p = p.next
	}

	// 'left' > max end, append to tail
	if p == nil {
		pre.next = &listNode{
			start: left,
			end:   right,
		}

		return
	}

	if p.start <= left && right <= p.end { // internal interval, skip
		return
	}

	newNode := &listNode{}
	if left < p.start { // p.start move to 'left'
		newNode.start = left
	} else {
		newNode.start = p.start
	}

	pre.next = newNode

	// find 'right' position
	for p != nil && right > p.end {
		p = p.next
	}

	if p == nil || right < p.start { // 'right' > max end and 'right' not in interval
		newNode.end = right
	} else { // in interval
		newNode.end = p.end
		p = p.next
	}

	newNode.next = p
}

func (r *RangeModule) QueryRange(left int, right int) bool {
	isContained := false
	for p := r.pre.next; p != nil; p = p.next {
		if p.start <= left && right <= p.end {
			isContained = true
			break
		}
	}

	return isContained
}

func (r *RangeModule) RemoveRange(left int, right int) {
	pre := r.pre
	p := pre.next

	// find 'left' position
	for p != nil && left > p.end {
		pre = p
		p = p.next
	}

	// 'left' > max end, del interval not exist
	if p == nil {
		return
	}

	// 'left' in interval, maintain [p.start, left)
	if left >= p.start {
		pre.next = &listNode{
			start: p.start,
			end:   left,
			next:  p,
		}

		pre = pre.next
	}

	// find 'right' position
	for p != nil && right >= p.end {
		p = p.next
	}

	if p != nil && right > p.start { // if 'right' > max end or 'right' not in interval, do nothing
		p.start = right
	}

	pre.next = p
}
