package mario

type GraphNode struct {
	Value    int
	Connects []int
}

func sumOfDistancesInTree(N int, edges [][]int) []int {
	if N == 1 {
		return []int{0}
	}

	// save nodes into a slice, index of slice have many uses ^_^
	nodeSlice := make([]*GraphNode, N) // N define on length is ok
	for i := range nodeSlice {
		nodeSlice[i] = &GraphNode{Value: i}
	}

	// make graph
	for _, e := range edges {
		nodeSlice[e[0]].Connects = append(nodeSlice[e[0]].Connects, e[1])
		nodeSlice[e[1]].Connects = append(nodeSlice[e[1]].Connects, e[0])
	}

	dist := make([]int, N)
	for i := 0; i < N; i++ {
		distI := make([]int, N)  // dist from i to other nodes
		flagI := make([]bool, N) // if a node has reached
		flagI[i] = true
		for _, v := range nodeSlice[i].Connects {
			distI[v] = 1
			flagI[v] = true
		}

		for count := 1; !isFinish(flagI); count++ {
			for k := range distI {
				if distI[k] == count {
					for _, v := range nodeSlice[k].Connects {
						if !flagI[v] {
							flagI[v] = true
							distI[v] = count + 1
						}
					}
				}
			}
		}

		// calc sum
		for _, v := range distI {
			dist[i] += v
		}
	}

	return dist
}

func isFinish(s []bool) bool {
	res := true
	for i := range s {
		if !s[i] {
			res = false
			break
		}
	}

	return res
}
