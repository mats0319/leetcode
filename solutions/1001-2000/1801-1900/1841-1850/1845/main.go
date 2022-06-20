package mario

type SeatManager struct {
	heap *binaryMinimumHeapOnArray
}

// Constructor all 'reserve' and 'unreserve' invoke are valid
func Constructor(n int) SeatManager {
	return SeatManager{
		heap: NewBinaryMinimumHeapOnArray(n),
	}
}

func (s *SeatManager) Reserve() int {
	v, _ := s.heap.Pop()

	return v
}

func (s *SeatManager) Unreserve(seatNumber int) {
	s.heap.Push(seatNumber)
}
