package mario

func findOrder(numCourses int, prerequisites [][]int) []int {
    res := make([]int, 0, numCourses)

    relationMap := makeRelationMap(prerequisites)
    visited := make([]bool, numCourses) // initialized
    count := numCourses

    for len(res) < numCourses {
        countBackup := count

        for i := 0; i < numCourses; i++ {
            if visited[i] {
                continue
            }

            v, ok := relationMap[i]
            if !ok || contains(res, v) {
                res = append(res, i)
                visited[i] = true
                count--
            }
        }

        if countBackup == count {
            break
        }
    }

    if len(res) < numCourses {
        res = nil
    }

    return res
}

// makeRelationMap return data relation in a map: value - pre-conditions
//   @param data: slice of [a, b], 'b' is pre-condition of 'a'
func makeRelationMap(data [][]int) map[int][]int {
    m := make(map[int][]int)

    for i := range data {
        relation := data[i]

        v, ok := m[relation[0]]
        if !ok {
            v = []int{relation[1]}
        } else {
            v = append(v, relation[1])
        }

        m[relation[0]] = v
    }

    return m
}

// contains return if 'sliceA' contains all items in 'sliceB', ignore orders
func contains(sliceA, sliceB []int) bool {
    isContained := true
    for i := range sliceB {
        isContainedInner := false
        for j := range sliceA {
            if sliceB[i] == sliceA[j] {
                isContainedInner = true
                break
            }
        }

        if !isContainedInner {
            isContained = false
            break
        }
    }

    return isContained
}
