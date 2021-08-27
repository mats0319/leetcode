package mario

type freqDoublyList struct {
    head *freqListNode
    tail *freqListNode
}

type freqListNode struct {
	frequency int
	data      *dataDoublyList
	prev      *freqListNode
	next      *freqListNode
}

func newFreqDoublyList() *freqDoublyList {
    head := &freqListNode{}
    tail := &freqListNode{}

    head.next = tail
    tail.prev = head

    return &freqDoublyList{
        head: head,
        tail: tail,
    }
}

func removeFreqNode(node *freqListNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (fln *freqListNode) addAfter(node *freqListNode) {
    fln.next.prev = node
    node.next = fln.next

    fln.next = node
    node.prev = fln
}
