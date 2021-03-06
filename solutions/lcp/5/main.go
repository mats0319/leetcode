package mario

func bonus(n int, leadership [][]int, operations [][]int) []int {
	team := make(map[int][]int, n)
	for _, relation := range leadership {
		team[relation[0]] = append(team[relation[0]], relation[1])
	}

	coins := make([]int, n+1) // index is matched with member

	res := make([]int, 0, len(operations))
	for _, op := range operations {
		switch op[0] {
		case 1:
			coins[op[1]] += op[2]
		case 2:
			subordinates := []int{op[1]}
			for len(subordinates) > 0 {
				nextLayer := make([]int, 0, n)
				for _, member := range subordinates {
					coins[member] += op[2]
					nextLayer = append(nextLayer, team[member]...)
				}

				subordinates = nextLayer
			}
		case 3:
			subordinates := []int{op[1]}
			sum := 0
			for len(subordinates) > 0 {
				nextLayer := make([]int, 0, n)
				for _, member := range subordinates {
					sum += coins[member]
					nextLayer = append(nextLayer, team[member]...)
				}

				subordinates = nextLayer
			}

			res = append(res, sum)
		}
	}

	return res
}
