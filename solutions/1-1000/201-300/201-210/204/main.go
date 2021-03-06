package mario

func countPrimes(n int) int {
	notPrime := make([]bool, n)

	count := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}

	return count
}
