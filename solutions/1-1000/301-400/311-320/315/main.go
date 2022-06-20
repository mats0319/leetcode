package mario

// countSmaller -10^4 <= nums[i] <= 10^4
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	list := newList()

	for i := len(res) - 1; i >= 0; i-- {
		res[i] = list.insert(nums[i])
	}

	return res
}

type node struct {
	value  int
	amount int

	next *node
}

func newList() *node {
	return &node{}
}

// insert add 'v' into list and return how many smaller values than 'v'
func (l *node) insert(v int) int {
	amount := 0

	pre := l
	for pre.next != nil && v > pre.next.value {
		amount += pre.next.amount
		pre = pre.next
	}

	if pre.next == nil {
		pre.next = &node{
			value:  v,
			amount: 1,
		}
	} else {
		if v == pre.next.value {
			pre.next.amount++
		} else { // v < pre.next.value
			pre.next = &node{
				value:  v,
				amount: 1,
				next:   pre.next,
			}
		}
	}

	return amount
}
