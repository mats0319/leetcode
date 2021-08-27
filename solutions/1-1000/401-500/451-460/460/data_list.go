package mario

type dataDoublyList struct {
	head *dataListNode
	tail *dataListNode
}

type dataListNode struct {
	Key       int
	Value     int
	Frequency int
	Prev      *dataListNode
	Next      *dataListNode
}

func newDataDoublyList() *dataDoublyList {
	head := &dataListNode{}
	tail := &dataListNode{}

	head.Next = tail
	tail.Prev = head

	return &dataDoublyList{
		head: head,
		tail: tail,
	}
}

func removeDataNode(node *dataListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (ddl *dataDoublyList) addToHead(node *dataListNode) {
	node.Next = ddl.head.Next
	ddl.head.Next.Prev = node

	node.Prev = ddl.head
	ddl.head.Next = node
}

func (ddl *dataDoublyList) removeTail() int {
	if ddl.isEmpty() {
		return -1
	}

	node := ddl.tail.Prev

	node.Prev.Next = ddl.tail
	ddl.tail.Prev = node.Prev

	return node.Key
}

func (ddl *dataDoublyList) isEmpty() (flag bool) {
	if ddl.head.Next == ddl.tail {
		flag = true
	}

	return
}
