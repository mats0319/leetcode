package mario

type keysList struct {
    head *keysNode
    tail *keysNode
}

type keysNode struct {
    key  string
    freq int
    prev *keysNode
    next *keysNode
}

func newKeysList() *keysList {
    headNode := &keysNode{}
    tailNode := &keysNode{}

    headNode.next = tailNode
    tailNode.prev = headNode

    return &keysList{
        head: headNode,
        tail: tailNode,
    }
}

// add insert node in front of list
func (kl *keysList) add(node *keysNode) {
    nextNode := kl.head.next

    kl.head.next = node
    node.next = nextNode

    nextNode.prev = node
    node.prev = kl.head
}

func (kl *keysList) isEmpty() bool {
    return kl.head.next == kl.tail
}

func removeKeyNode(node *keysNode) {
    prevNode := node.prev
    nextNode := node.next

    prevNode.next = nextNode
    nextNode.prev = prevNode
}
