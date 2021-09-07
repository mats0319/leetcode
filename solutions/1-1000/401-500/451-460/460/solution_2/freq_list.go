package mario

type freqList struct {
	head *freqNode
	tail *freqNode
}

type freqNode struct {
	frequency int
	data      *cacheList
	prev      *freqNode
	next      *freqNode
}

func newFreqList() *freqList {
	headNode := &freqNode{}
	tailNode := &freqNode{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &freqList{
		head: headNode,
		tail: tailNode,
	}
}

func (f *freqNode) remove() {
	f.prev.next = f.next
	f.next.prev = f.prev
}

func (f *freqNode) addBehind(node *freqNode) {
	next := f.next

	f.next = node
	node.next = next

	next.prev = node
	node.prev = f
}

func (fl *freqList) addToHead(node *freqNode) {
	next := fl.head.next

	fl.head.next = node
	node.next = next

	next.prev = node
	node.prev = fl.head
}
