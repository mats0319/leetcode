package mario

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))

	list := make([]int, 0, len(isConnected))
	result := 0
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			visited[i] = true
			list = append(list, i)
			result++

			for len(list) > 0 {
				item := list[0]
				list = list[1:]

				for j := 0; j < len(isConnected[item]); j++ {
					if isConnected[item][j] == 1 && !visited[j] {
						list = append(list, j)
						visited[j] = true
					}
				}
			}
		}
	}

	return result
}
