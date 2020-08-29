package mario

// [0, 1]: 1 -> 0
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges    = make([][]int, numCourses)
		indegree = make([]int, numCourses)
		bfsRes   = make([]int, 0)
	)

	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		indegree[v[0]]++
	}

	for i := range indegree {
		if indegree[i] == 0 {
			bfsRes = append(bfsRes, i)
		}
	}

	if len(bfsRes) <= 0 {
		return false
	}

	for i := 0; i < len(bfsRes); i++ {
		for j := 0; j < len(edges[bfsRes[i]]); j++ {
			indegree[edges[bfsRes[i]][j]]--
			if indegree[edges[bfsRes[i]][j]] == 0 {
				bfsRes = append(bfsRes, edges[bfsRes[i]][j])
			}
		}
	}

	return len(bfsRes) == numCourses
}
