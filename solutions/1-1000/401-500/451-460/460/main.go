package mario

type LFUCache struct {
	recent   map[int]*listNode // freq - first node matched freq
	count    map[int]int       // freq - amount
	cache    map[int]*listNode // key - node
	list     *doublyList
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		recent:   make(map[int]*listNode, capacity),
		count:    make(map[int]int),
		cache:    make(map[int]*listNode, capacity),
		list:     newDoublyList(),
		capacity: capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}

	node, ok := lfu.cache[key]
	if !ok { // key不存在
		return -1
	}

	// key已存在
	next := node.next

	if lfu.count[node.frequency+1] > 0 {
		// 存在其他使用次数为n+1的缓存，将指定缓存移动到所有使用次数为n+1的节点之前
		removeNode(node)
		addBefore(lfu.recent[node.frequency+1], node)
	} else if lfu.count[node.frequency] > 1 && lfu.recent[node.frequency] != node {
		// 不存在其他使用次数为n+1的缓存，但存在其他使用次数为n的缓存，且当前节点不是最近的节点
		// 将指定缓存移动到所有使用次数为n的节点之前
		removeNode(node)
		addBefore(lfu.recent[node.frequency], node)
	}

	// 更新recent
	lfu.recent[node.frequency+1] = node
	if lfu.count[node.frequency] <= 1 { // 不存在其他freq = n的节点，recent置空
		lfu.recent[node.frequency] = nil
	} else if lfu.recent[node.frequency] == node { // 存在其他freq = n的节点，且recent = node，将recent向后移动一位
		lfu.recent[node.frequency] = next
	}

	// 更新使用次数对应的节点数
	lfu.count[node.frequency+1]++
	lfu.count[node.frequency]--

	// 更新缓存使用次数
	node.frequency++

	return node.value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	node, ok := lfu.cache[key]
	if ok { // key已存在
		lfu.Get(key)
		node.value = value

		return
	}

	// key不存在
	if len(lfu.cache) >= lfu.capacity { // 缓存已满，删除最后一个节点，相应更新cache、count、recent（条件）
		tailNode := lfu.list.tail.prev

		lfu.list.removeTail()

		if lfu.count[tailNode.frequency] <= 1 {
			lfu.recent[tailNode.frequency] = nil
		}
		lfu.count[tailNode.frequency]--
		delete(lfu.cache, tailNode.key)
	}

	newNode := &listNode{
		key:       key,
		value:     value,
		frequency: 1,
	}

	// 插入新的缓存节点
	if lfu.count[1] > 0 {
		addBefore(lfu.recent[1], newNode)
	} else {
		addBefore(lfu.list.tail, newNode)
	}

	// 更新recent、count、cache
	lfu.recent[1] = newNode
	lfu.count[1]++
	lfu.cache[key] = newNode
}
