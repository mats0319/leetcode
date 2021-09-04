package mario

type LRUCache struct {
	cache    map[int]*doublyListNode
	head     *doublyListNode // dummy node
	tail     *doublyListNode // dummy node
	capacity int
}

type doublyListNode struct {
	key   int
	value int
	prev  *doublyListNode
	next  *doublyListNode
}

func Constructor(capacity int) LRUCache {
	dl := newDoublyList()

	return LRUCache{
		cache:    make(map[int]*doublyListNode, capacity),
		head:     dl,
		tail:     dl.next,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}

	removeNode(v)
	this.head.addToHead(v)

	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.value = value
		this.cache[key] = v

		removeNode(v)
		this.head.addToHead(v)

		return
	}

	if len(this.cache) >= this.capacity {
		node := this.tail.prev
		removeNode(node)
		delete(this.cache, node.key)
	}

	newNode := &doublyListNode{
		key:   key,
		value: value,
	}

	this.cache[key] = newNode
	this.head.addToHead(newNode)
}

func newDoublyList() *doublyListNode {
	headNode := &doublyListNode{}
	tailNode := &doublyListNode{}
	headNode.next = tailNode
	tailNode.prev = headNode

	return headNode
}

func removeNode(node *doublyListNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (dl *doublyListNode) addToHead(node *doublyListNode) {
	dl.next.prev = node
	node.next = dl.next

	dl.next = node
	node.prev = dl
}
