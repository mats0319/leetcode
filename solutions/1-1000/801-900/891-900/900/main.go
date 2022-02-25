package mario

type RLEIterator struct {
	data  []int
	index int
}

func Constructor(encoding []int) RLEIterator {
	data := make([]int, 0)

	for i := 0; i < len(encoding); i += 2 {
		for j := 0; j < encoding[i]; j++ {
			data = append(data, encoding[i+1])
		}
	}

	return RLEIterator{
		data: data,
	}
}

func (i *RLEIterator) Next(n int) int {
    i.index += n

    if i.index > len(i.data) {
        return -1
    }

    return i.data[i.index-1]
}
