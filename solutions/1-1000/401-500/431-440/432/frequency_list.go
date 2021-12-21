package mario

type frequencyList struct {
	head *frequencyNode
	tail *frequencyNode
}

type frequencyNode struct {
	freq    int
	keyList *keysList
	prev    *frequencyNode
	next    *frequencyNode
}

func newFrequencyList() *frequencyList {
	headNode := &frequencyNode{keyList: newKeysList()}
	tailNode := &frequencyNode{keyList: newKeysList()}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &frequencyList{
		head: headNode,
		tail: tailNode,
	}
}

func (fl *frequencyList) isEmpty() bool {
	return fl.head.next == fl.tail
}

func addAfterFreqNode(currNode, newNode *frequencyNode) {
	nextNode := currNode.next

	currNode.next = newNode
	newNode.next = nextNode

	nextNode.prev = newNode
	newNode.prev = currNode
}

func removeFreqNode(node *frequencyNode) {
	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
}
