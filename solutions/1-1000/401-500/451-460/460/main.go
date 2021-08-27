package mario

type LFUCache struct {
	freq     map[int]*freqListNode // freq - node
	data     map[int]*dataListNode // key - node
	list     *freqDoublyList
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freq:     make(map[int]*freqListNode),
		data:     make(map[int]*dataListNode),
		list:     newFreqDoublyList(),
		capacity: capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	dataNode, ok := lfu.data[key]
	if !ok {
		return -1
	}

	// remove data node
	freq := dataNode.Frequency
	removeDataNode(dataNode)
	freqNode, _ := lfu.freq[freq] // avoid panic
	if freqNode.data.isEmpty() {
		prev := freqNode.prev

		removeFreqNode(freqNode)
		delete(lfu.freq, freqNode.frequency)

		freqNode = prev
	}

	// add data node
	dataNode.Frequency++
	nextFreqNode, ok := lfu.freq[freq+1]
	if !ok {
		nextFreqNode = &freqListNode{
			frequency: 1,
			data:      newDataDoublyList(),
		}

		freqNode.addAfter(nextFreqNode)
		lfu.freq[1] = nextFreqNode
	}
	nextFreqNode.data.addToHead(dataNode)

	return dataNode.Value
}

func (lfu *LFUCache) Put(key int, value int) {
	if _, ok := lfu.data[key]; ok {
		_ = lfu.Get(key) // updated dataNode.freq

		dataNode, _ := lfu.data[key]
		dataNode.Value = value
		lfu.data[key] = dataNode

		return
	}

	// judge if it reach capacity
	if len(lfu.data) >= lfu.capacity {
		freqNode := lfu.list.head.next
		k := freqNode.data.removeTail()
		delete(lfu.data, k)
		if freqNode.data.isEmpty() {
			removeFreqNode(freqNode)
			delete(lfu.freq, freqNode.frequency)
		}
	}

	freqNode := lfu.list.head.next
	if freqNode.frequency != 1 {
		freqNode = &freqListNode{
			frequency: 1,
			data:      newDataDoublyList(),
		}

		lfu.list.head.addAfter(freqNode)
		lfu.freq[1] = freqNode
	}

	dataNode := &dataListNode{
		Key:       key,
		Value:     value,
		Frequency: 1,
	}
	freqNode.data.addToHead(dataNode)
	lfu.data[key] = dataNode

	return
}
