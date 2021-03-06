package mario

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		count := 0
		for x := i; x > 0; x &= x - 1 {
			count++
		}

		res[i] = count
	}

	return res
}
