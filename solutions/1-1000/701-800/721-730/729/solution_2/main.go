package mario

type MyCalendar struct {
	list *listNode
}

type listNode struct {
	start int
	end   int
	next  *listNode
}

func Constructor() MyCalendar {
	return MyCalendar{
		list: &listNode{},
	}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	pre := mc.list
	for pre.next != nil && pre.next.start <= start {
		pre = pre.next
	}

	if start < pre.end || (pre.next != nil && end > pre.next.start) {
		return false
	}

	if pre.end == start {
		pre.end = end
	} else { // pre.end < start
		newNode := &listNode{
			start: start,
			end:   end,
			next:  pre.next,
		}

		pre.next = newNode
		pre = pre.next
	}

	if pre.next != nil && end == pre.next.start {
		pre.end = pre.next.end

		pre.next = pre.next.next
	}

	return true
}
