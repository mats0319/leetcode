package mario

type RecentCounter struct {
	list   *listNode
	tail   *listNode
	length int
}

type listNode struct {
	time int
	next *listNode
}

func Constructor() RecentCounter {
	preNode := &listNode{}

	return RecentCounter{
		list: preNode,
		tail: preNode,
	}
}

// Ping param t always bigger than last time
func (rc *RecentCounter) Ping(t int) int {
    node := &listNode{
        time: t,
    }

    rc.tail.next = node
    rc.tail = rc.tail.next
    rc.length++

    for rc.list.next != nil && rc.list.next.time + 3000 < t {
        rc.list = rc.list.next
        rc.length--
    }

    return rc.length
}
