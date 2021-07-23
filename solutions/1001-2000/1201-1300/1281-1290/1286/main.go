package mario

type CombinationIterator struct {
	str   string
	index []int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	arr := make([]int, combinationLength)
	for i := range arr {
		arr[i] = i
	}

	return CombinationIterator{
		str:   characters,
		index: arr,
	}
}

func (ci *CombinationIterator) Next() string {
	res := make([]byte, 0, len(ci.index))
	for _, v := range ci.index {
		res = append(res, ci.str[v])
	}

	ci.nextIndex()

	return string(res)
}

func (ci *CombinationIterator) HasNext() bool {
	return ci.index[0]+len(ci.index) <= len(ci.str)
}

func (ci *CombinationIterator) nextIndex() {
	if !ci.HasNext() {
		return
	}

	lastIndex := len(ci.index) - 1
	if ci.index[lastIndex] < len(ci.str)-1 {
		ci.index[lastIndex]++
		return
	}

	index := lastIndex
	for count := 1; index >= 0 && ci.index[index]+count >= len(ci.str); index-- {
		count++
	}

	if index < 0 {
		ci.index[0] = len(ci.str)
		return
	}

	ci.index[index]++
	for i := index + 1; i < len(ci.index); i++ {
		ci.index[i] = ci.index[i-1] + 1
	}
}
