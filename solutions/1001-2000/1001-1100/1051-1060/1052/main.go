package mario

func maxSatisfied(customers []int, grumpy []int, X int) int {
	res := 0
	for i := range customers {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}

	increment := 0
	for i := 0; i < X; i++ {
		increment += grumpy[i] * customers[i]
	}

	maxIncrement := increment
	for i := 1; i+X-1 < len(customers); i++ {
		increment -= grumpy[i-1] * customers[i-1]
		increment += grumpy[i+X-1] * customers[i+X-1]

		if maxIncrement < increment {
		    maxIncrement = increment
        }
	}

	return maxIncrement + res
}
