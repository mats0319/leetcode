package mario

type OrderedStream struct {
	data  []string
	index int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		data:  make([]string, n+1),
		index: 1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.data[idKey] = value

	if idKey != this.index {
		return nil
	}

	res := make([]string, 0, len(this.data)-idKey+1)
	for this.index < len(this.data) && len(this.data[this.index]) > 0 {
		res = append(res, this.data[this.index])
		this.index++
	}

	return res
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
