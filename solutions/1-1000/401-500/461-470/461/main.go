package mario

func hammingDistance(x int, y int) int {
	num := x ^ y

	count := 0
	for num > 0 {
		if num%2 == 1 {
			count++
		}

		num >>= 1
	}

	return count
}
