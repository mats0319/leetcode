package mario

type SummaryRanges struct {
	list *listNode // list with pre-node
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		list: &listNode{},
	}
}

func (s *SummaryRanges) AddNum(val int) {
	pre := s.list
	for pre.next != nil {
		if pre.next.start <= val && val <= pre.next.end {
			return // ignore duplicated val
		}

		if val < pre.next.start {
			break
		}

		pre = pre.next
	}

	if pre == s.list {
		if pre.next != nil && val+1 == pre.next.start {
			pre.next.start = val
		} else {
			nextNode := &listNode{
				start: val,
				end:   val,
			}

			nextNode.next = pre.next
			pre.next = nextNode
		}
	} else if pre.next == nil {
		if pre.end+1 == val {
			pre.end = val
		} else {
			pre.next = &listNode{
				start: val,
				end:   val,
			}
		}
	} else {
		if pre.end+1 == val && val+1 == pre.next.start {
			pre.end = pre.next.end
			pre.next = pre.next.next
		} else if pre.end+1 == val {
			pre.end = val
		} else if val+1 == pre.next.start {
			pre.next.start = val
		} else {
			nextNode := &listNode{
				start: val,
				end:   val,
			}

			nextNode.next = pre.next
			pre.next = nextNode
		}
	}
}

func (s *SummaryRanges) GetIntervals() [][]int {
	res := make([][]int, 0)
	for p := s.list.next; p != nil; p = p.next {
		res = append(res, []int{p.start, p.end})
	}

	return res
}

type listNode struct {
	start int
	end   int
	next  *listNode
}
