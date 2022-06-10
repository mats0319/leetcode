package mario

type RLEIterator struct {
	count []int
	value []int
	index int
}

func Constructor(encoding []int) RLEIterator {
	ins := RLEIterator{
		count: make([]int, 0, len(encoding)/2),
		value: make([]int, 0, len(encoding)/2),
	}

	for i := 0; i < len(encoding); i += 2 {
		if encoding[i] > 0 {
			ins.count = append(ins.count, encoding[i])
			ins.value = append(ins.value, encoding[i+1])
		}
	}

	return ins
}

func (ins *RLEIterator) Next(n int) int {
	for ins.index < len(ins.count) && n > ins.count[ins.index] {
		n -= ins.count[ins.index]
		ins.index++
	}

	if ins.index >= len(ins.count) {
		return -1
	}

	value := ins.value[ins.index]
	if n == ins.count[ins.index] {
		ins.index++
	} else {
		ins.count[ins.index] -= n
	}

	return value
}
