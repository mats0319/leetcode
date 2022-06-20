package mario

type TimeMap struct {
	data map[string]*timeValueNode // pre node
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string]*timeValueNode),
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	timeValueList, ok := m.data[key]
	if !ok { // key is not exist
		timeValueList = &timeValueNode{}
	}

	timeValueList.insert(timestamp, value)

	m.data[key] = timeValueList
}

func (m *TimeMap) Get(key string, timestamp int) string {
	timeValueList, ok := m.data[key]
	if !ok { // key is not exist
		return ""
	}

	pre := timeValueList
	for pre.next != nil && pre.next.timestamp <= timestamp {
		pre = pre.next
	}

	return pre.value
}

type timeValueNode struct {
	timestamp int
	value     string
	next      *timeValueNode
}

func (n *timeValueNode) insert(t int, v string) {
	pre := n
	p := n.next
	for p != nil && p.timestamp < t {
		pre = pre.next
		p = p.next
	}

	pre.next = &timeValueNode{
		timestamp: t,
		value:     v,
		next:      p,
	}
}
