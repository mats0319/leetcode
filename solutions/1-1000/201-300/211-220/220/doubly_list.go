package mario

type doublyList struct {
	head    *listNode
	tail    *listNode
	nodeMap map[int]*listNode // index - node
}

type listNode struct {
	index int
	value int
	prev  *listNode
	next  *listNode
}

func newDoublyList() *doublyList {
	headNode := &listNode{}
	tailNode := &listNode{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &doublyList{
		head:    headNode,
		tail:    tailNode,
		nodeMap: make(map[int]*listNode),
	}
}

func (dl *doublyList) addReturn(index, value int) *listNode {
	newNode := &listNode{
		index: index,
		value: value,
	}

	dl.nodeMap[index] = newNode

	pre := dl.head
	for pre.next != dl.tail && pre.next.value < newNode.value {
		pre = pre.next
	}

	next := pre.next

	pre.next = newNode
	newNode.next = next

	next.prev = newNode
	newNode.prev = pre

	return newNode
}

func (dl *doublyList) del(index int) {
	node, _ := dl.nodeMap[index]
	delete(dl.nodeMap, index)

	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (dl *doublyList) validDifference(node *listNode, target int) bool {
	difference := -1

	if node.prev != dl.head {
		difference = node.value - node.prev.value
		if difference < 0 {
			difference *= -1
		}
	}

	if node.next != dl.tail {
		nextDifference := node.value - node.next.value
		if nextDifference < 0 {
			nextDifference *= -1
		}

		if difference < 0 || difference > nextDifference {
			difference = nextDifference
		}
	}

	return 0 <= difference && difference <= target
}
