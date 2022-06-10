package mario

type Bitset struct {
	fixed   []bool
	count   int
	flipped bool
}

func Constructor(size int) Bitset {
	return Bitset{
		fixed: make([]bool, size),
	}
}

func (s *Bitset) Fix(idx int) {
	if s.fixed[idx] == s.flipped {
		s.fixed[idx] = !s.fixed[idx]
        s.count++
	}
}

func (s *Bitset) Unfix(idx int) {
	if s.fixed[idx] != s.flipped {
		s.fixed[idx] = !s.fixed[idx]
        s.count--
	}
}

func (s *Bitset) Flip() {
	s.flipped = !s.flipped
	s.count = len(s.fixed) - s.count
}

func (s *Bitset) All() bool {
	return s.count == len(s.fixed)
}

func (s *Bitset) One() bool {
	return s.count > 0
}

func (s *Bitset) Count() int {
	return s.count
}

func (s *Bitset) ToString() string {
	byteSlice := make([]byte, 0, len(s.fixed))
	for i := range s.fixed {
		if s.fixed[i] == s.flipped {
			byteSlice = append(byteSlice, '0')
		} else {
			byteSlice = append(byteSlice, '1')
		}
	}

	return string(byteSlice)
}
