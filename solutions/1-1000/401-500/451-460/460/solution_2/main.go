package mario

type LFUCache struct {
	freqList  *freqList
	freqMap   map[int]*freqNode // freq - node
	cacheList *cacheList
	cacheMap  map[int]*cacheNode // key - node
	capacity  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freqList:  newFreqList(),
		freqMap:   make(map[int]*freqNode),
		cacheList: newCacheList(),
		cacheMap:  make(map[int]*cacheNode),
		capacity:  capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}

	node, ok := lfu.cacheMap[key]
	if !ok { // key 不存在
		return -1
	}

	// key 存在
	node.remove()

	freq := node.frequency
	fNode := lfu.freqMap[freq]

	newFreqNode, ok := lfu.freqMap[freq+1]
	if !ok { // 不存在 freq = n+1 的节点
		newFreqNode = &freqNode{
			frequency: freq + 1,
			data:      newCacheList(),
		}

		fNode.addBehind(newFreqNode)
		lfu.freqMap[freq+1] = newFreqNode
	}

	node.frequency++
	newFreqNode.data.addToHead(node)

	if fNode.data.isEmpty() { // 频率节点(freq = n)为空，后删除是为了帮助freq = n+1的频率节点在新增时定位
		fNode.remove()
		delete(lfu.freqMap, freq)
	}

	return node.value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	node, ok := lfu.cacheMap[key]
	if ok {
		lfu.Get(key)
		node.value = value

		return
	}

	if len(lfu.cacheMap) >= lfu.capacity {
		fNode := lfu.freqList.head.next

		delNode := fNode.data.tail.prev
		delNode.remove()
		delete(lfu.cacheMap, delNode.key)

		if fNode.data.isEmpty() && fNode.frequency > 1 {
			delete(lfu.freqMap, fNode.frequency)
		}
	}

	fNode, ok := lfu.freqMap[1]
	if !ok {
		fNode = &freqNode{
			frequency: 1,
			data:      newCacheList(),
		}

		lfu.freqList.addToHead(fNode)
		lfu.freqMap[1] = fNode
	}

	newCacheNode := &cacheNode{
		key:       key,
		value:     value,
		frequency: 1,
	}

	fNode.data.addToHead(newCacheNode)
	lfu.cacheMap[key] = newCacheNode
}
