package mario

// n >= 1
func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res ^= start + i*2
	}

	return res
}
