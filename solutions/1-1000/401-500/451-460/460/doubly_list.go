package mario

type doublyList struct {
	head *listNode
	tail *listNode
}

type listNode struct {
	key       int
	value     int
	frequency int
	prev      *listNode
	next      *listNode
}

func newDoublyList() *doublyList {
	head := &listNode{}
	tail := &listNode{}

	head.next = tail
	tail.prev = head

	return &doublyList{
		head: head,
		tail: tail,
	}
}

func removeNode(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func addBefore(currNode *listNode, newNode *listNode) {
	pre := currNode.prev

	pre.next = newNode
	newNode.next = currNode

	currNode.prev = newNode
	newNode.prev = pre
}

func (dl *doublyList) removeTail() {
	pre := dl.tail.prev.prev

	pre.next = dl.tail
	dl.tail.prev = pre
}

func (dl *doublyList) isEmpty() bool {
	return dl.head.next == dl.tail
}
