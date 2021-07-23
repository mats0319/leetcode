package mario

type LFUCache struct {
	data      map[int]*dataListNode      // key - node length: len(map)
	freqNode  map[int]*frequencyListNode // frequency - node
	frequency *frequencyList
	capacity  int
}

type frequencyList struct {
	head *frequencyListNode
	tail *frequencyListNode
}

type frequencyListNode struct {
	frequency int
	length    int
	data      *dataList
	prev      *frequencyListNode
	next      *frequencyListNode
}

type dataList struct {
	head *dataListNode
	tail *dataListNode
}

type dataListNode struct {
	key       int
	value     int
	frequency int
	prev      *dataListNode
	next      *dataListNode
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		data:      make(map[int]*dataListNode, capacity),
		freqNode:  make(map[int]*frequencyListNode),
		frequency: newFrequencyList(),
		capacity:  capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	node, ok := lfu.data[key]
	if !ok {
		return -1
	}

	node.remove()
	freq := node.frequency
	freqNode := lfu.freqNode[freq]
	freqNode.length--
	if freqNode.length <= 0 {
		freqNode.remove()
		delete(lfu.freqNode, freqNode.frequency)
	}

	node.frequency++
	lfu.data[key] = node

	if freqNode.next.frequency != freq+1 {
		freqNode.addNext(freq+1)
	}
	next := freqNode.next
	next.data.addToHead(node)
	next.length++

	return node.value
}

func (lfu *LFUCache) Put(key int, value int) {
	if node, ok := lfu.data[key]; ok {
		// update data and move node
		node.remove()
		freq := node.frequency
		freqNode := lfu.freqNode[freq]
		freqNode.length--
		if freqNode.length <= 0 {
			freqNode.remove()
			delete(lfu.freqNode, freqNode.frequency)
		}

		node.value = value
		node.frequency++
		lfu.data[key] = node

		if freqNode.next.frequency != freq+1 {
			freqNode.addNext(freq+1)
		}
		next := freqNode.next
		next.data.addToHead(node)
		next.length++

		return
	}

	node := &dataListNode{
		key:       key,
		value:     value,
		frequency: 1,
	}

	lfu.data[key] = node
	if len(lfu.data) > lfu.capacity {
		freqNode := lfu.frequency.head.next
		k := freqNode.data.removeTail()
		delete(lfu.data, k)

		freqNode.length--
		if freqNode.length <= 0 {
			freqNode.remove()
			delete(lfu.freqNode, freqNode.frequency)
		}
	}

	if lfu.frequency.head.next.frequency != 1 {
		freqNode := &frequencyListNode{
			frequency: 1,
			length:    0,
			data:      newDataList(),
		}

		lfu.freqNode[1] = freqNode

		lfu.frequency.addToHead(freqNode)
	}

	freqNode, _ := lfu.freqNode[1]
	freqNode.data.addToHead(node)
	freqNode.length++

	return
}

func newFrequencyList() *frequencyList {
	headNode := &frequencyListNode{}
	tailNode := &frequencyListNode{}
	headNode.next = tailNode
	tailNode.prev = headNode

	return &frequencyList{
		head: headNode,
		tail: tailNode,
	}
}

func (fl *frequencyList) addToHead(node *frequencyListNode) {
	fl.head.next.prev = node
	node.next = fl.head.next

	fl.head.next = node
	node.prev = fl.head
}

func (fln *frequencyListNode) addNext(frequency int) {
	node := &frequencyListNode{
		frequency: frequency,
		length:    0,
		data:      newDataList(),
		prev:      fln,
		next:      fln.next,
	}

	fln.next.prev = node
	fln.next = node
}

// remove itself
func (fln *frequencyListNode) remove() {
	fln.prev.next = fln.next
	fln.next.prev = fln.prev
}

func newDataList() *dataList {
	headNode := &dataListNode{}
	tailNode := &dataListNode{}
	headNode.next = tailNode
	tailNode.prev = headNode

	return &dataList{
		head: headNode,
		tail: tailNode,
	}
}

// make sure dl is not nil
func (dl *dataList) addToHead(node *dataListNode) {
	dl.head.next.prev = node
	node.next = dl.head.next

	dl.head.next = node
	node.prev = dl.head
}

func (dl *dataList) removeTail() int {
	node := dl.tail.prev

	node.prev.next = dl.tail
	dl.tail.prev = node.prev

	return node.key
}

func (dl *dataListNode) remove() {
	dl.prev.next = dl.next
	dl.next.prev = dl.prev
}
