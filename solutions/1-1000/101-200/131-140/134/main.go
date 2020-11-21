package mario

// gas.length == cost.length
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	res := -1

	gasLeft := 0
	distance := 0
	for i := 0; i < length*2; i++ {
		gasLeft += gas[i%length] - cost[i%length]
		if gasLeft < 0 {
			gasLeft = 0
			distance = 0
			continue
		}

		distance++
		if distance >= length {
			res = (i + 1) % length
			break
		}
	}

	return res
}
