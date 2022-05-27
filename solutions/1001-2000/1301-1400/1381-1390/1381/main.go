package mario

type CustomStack struct {
	data     []int
	top      int
	capacity int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		data:     make([]int, maxSize),
		top:      0,
		capacity: maxSize,
	}
}

func (s *CustomStack) Push(x int) {
	if s.top >= s.capacity {
		return
	}

	s.data[s.top] = x
	s.top++
}

func (s *CustomStack) Pop() int {
	if s.top < 1 {
		return -1
	}

	s.top--

	return s.data[s.top]
}

func (s *CustomStack) Increment(k int, val int) {
	k = small(k, s.top) // k-1 is index

	for i := 0; i < k; i++ {
		s.data[i] += val
	}
}

func small(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
