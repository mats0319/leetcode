package mario

type cacheList struct {
	head *cacheNode
	tail *cacheNode
}

type cacheNode struct {
	key       int
	value     int
	frequency int
	prev      *cacheNode
	next      *cacheNode
}

func newCacheList() *cacheList {
	headNode := &cacheNode{}
	tailNode := &cacheNode{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &cacheList{
		head: headNode,
		tail: tailNode,
	}
}

func (c *cacheNode) remove() {
	c.prev.next = c.next
	c.next.prev = c.prev
}

func (cl *cacheList) addToHead(node *cacheNode) {
	next := cl.head.next

	cl.head.next = node
	node.next = next

	next.prev = node
	node.prev = cl.head
}

func (cl *cacheList) isEmpty() bool {
	return cl.head.next == cl.tail
}
