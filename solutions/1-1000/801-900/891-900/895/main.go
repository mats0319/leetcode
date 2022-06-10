package mario

type FreqStack struct {
	// use freq list and value list, with some map for O(1)
	// if push, keep node in old freq, insert new node at first of new freq
	// if pop, just del first value node of first freq node
	// map: value - freq, for quick find position of value
	// map: freq - freq node
	// map: value - data node, because not need to del data node, this map is unnecessary
	list         *freqDoublyList   // pre
	valueFreqMap map[int]int       // value - freq
	freqNodeMap  map[int]*freqNode // freq - freq node
}

func Constructor() FreqStack {
	return FreqStack{
		list:         newFreqDoublyList(),
		valueFreqMap: make(map[int]int),
		freqNodeMap:  make(map[int]*freqNode),
	}
}

func (s *FreqStack) Push(val int) {
	valueFreq, ok := s.valueFreqMap[val]
	if !ok { // new 'value'
		freq1Node, ok := s.freqNodeMap[1]
		if !ok {
			freq1Node = &freqNode{
				freq:       1,
				data:       &dataNode{},
				dataLength: 0,
			}

			s.freqNodeMap[1] = freq1Node

			// insert freqNode:1
			prev := s.list.tail.prev
			next := s.list.tail

			prev.next = freq1Node
			freq1Node.next = next

			next.prev = freq1Node
			freq1Node.prev = prev
		}

		s.valueFreqMap[val] = 1

		// insert dataNode:val
		freq1Node.data.next = &dataNode{
			value: val,
			next:  freq1Node.data.next,
		}
		freq1Node.dataLength++

		return
	}

	oldFreqNode, _ := s.freqNodeMap[valueFreq] // must exist

	newFreqNode := oldFreqNode.prev
	if newFreqNode.freq != valueFreq+1 { // new 'freq node' not exist
		newFreqNode = &freqNode{
			freq:       oldFreqNode.freq + 1,
			data:       &dataNode{},
			dataLength: 0,
		}

		s.freqNodeMap[newFreqNode.freq] = newFreqNode

		prev := oldFreqNode.prev
		next := oldFreqNode

		prev.next = newFreqNode
		newFreqNode.next = next

		next.prev = newFreqNode
		newFreqNode.prev = prev
	}

	s.valueFreqMap[val] = valueFreq + 1

	newFreqNode.data.next = &dataNode{
		value: val,
		next:  newFreqNode.data.next,
	}
	newFreqNode.dataLength++
}

// Pop always valid
func (s *FreqStack) Pop() int {
	firstFreqNode := s.list.head.next

	delNode := firstFreqNode.data.next

	firstFreqNode.data.next = delNode.next
	firstFreqNode.dataLength--
	if firstFreqNode.dataLength <= 0 { // del freq node
		delete(s.freqNodeMap, firstFreqNode.freq)

		prev := s.list.head
		next := firstFreqNode.next

		prev.next = next
		next.prev = prev
	}

	if firstFreqNode.freq-1 > 0 {
		s.valueFreqMap[delNode.value] = firstFreqNode.freq - 1
	} else {
		delete(s.valueFreqMap, delNode.value)
	}

	return delNode.value
}

type freqDoublyList struct {
	head *freqNode
	tail *freqNode
}

func newFreqDoublyList() *freqDoublyList {
	headNode := &freqNode{}
	tailNode := &freqNode{}

	headNode.next = tailNode
	tailNode.prev = headNode

	return &freqDoublyList{
		head: headNode,
		tail: tailNode,
	}
}

type freqNode struct {
	freq int

	data       *dataNode // pre
	dataLength int

	prev *freqNode
	next *freqNode
}

type dataNode struct {
	value int
	next  *dataNode
}
