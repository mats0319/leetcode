package mario

type AllOne struct {
	list        *frequencyList         // freq list
	freqNodeMap map[int]*frequencyNode // freq - freq node
	keyNodeMap  map[string]*keysNode   // key - key node
}

func Constructor() AllOne {
	return AllOne{
		list:        newFrequencyList(),
		freqNodeMap: make(map[int]*frequencyNode),
		keyNodeMap:  make(map[string]*keysNode),
	}
}

func (a *AllOne) Inc(key string) {
	prevFreqNode := a.list.head

	keyNode, ok := a.keyNodeMap[key]
	if !ok {
		keyNode = &keysNode{
			key:  key,
			freq: 1,
		}

		a.keyNodeMap[key] = keyNode
	} else {
		currFreqNode, _ := a.freqNodeMap[keyNode.freq] // we consider that if keyNode is exist, it's freqNode will exist
		prevFreqNode = currFreqNode

        // rm keyNode from currFreqNode and maintain currFreqNode
		removeKeyNode(keyNode)
		if currFreqNode.keyList.isEmpty() {
			prevFreqNode = prevFreqNode.prev

			removeFreqNode(currFreqNode)
			delete(a.freqNodeMap, keyNode.freq)
		}

		keyNode.freq++
	}

    // add keyNode to freqNode:freq=keyNode.freq.keyList
	nextFreqNode, ok := a.freqNodeMap[keyNode.freq]
	if !ok {
		nextFreqNode = &frequencyNode{
			freq:    keyNode.freq,
			keyList: newKeysList(),
		}

		a.freqNodeMap[keyNode.freq] = nextFreqNode
		addAfterFreqNode(prevFreqNode, nextFreqNode)
	}

	nextFreqNode.keyList.add(keyNode)
}

func (a *AllOne) Dec(key string) {
	keyNode, ok := a.keyNodeMap[key]
	if !ok {
		return
	}

	currFreqNode, _ := a.freqNodeMap[keyNode.freq]
    prevFreqNode := currFreqNode.prev

	// rm keyNode from currFreqNode and maintain currFreqNode
	removeKeyNode(keyNode)
	if currFreqNode.keyList.isEmpty() {
		removeFreqNode(currFreqNode)
		delete(a.freqNodeMap, currFreqNode.freq)
	}

	// freq = 1, del keyNode
	if keyNode.freq <= 1 {
		delete(a.keyNodeMap, key)
		return
	}

	// freq > 1, add keyNode to freqNode:freq=keyNode.freq-1.keyList
	keyNode.freq--

	// make sure freqNode:freq=keyNode.freq-1 is exist
	freqNodeSubOne, ok := a.freqNodeMap[keyNode.freq]
	if !ok {
		freqNodeSubOne = &frequencyNode{
			freq:    keyNode.freq,
			keyList: newKeysList(),
		}

		a.freqNodeMap[keyNode.freq] = freqNodeSubOne
		addAfterFreqNode(prevFreqNode, freqNodeSubOne)
	}

	freqNodeSubOne.keyList.add(keyNode)
}

func (a *AllOne) GetMaxKey() string {
	if a.list.isEmpty() {
		return ""
	}

	return a.list.tail.prev.keyList.head.next.key
}

func (a *AllOne) GetMinKey() string {
	if a.list.isEmpty() {
		return ""
	}

	return a.list.head.next.keyList.head.next.key
}
