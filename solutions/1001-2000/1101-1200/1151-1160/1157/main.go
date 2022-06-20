package mario

type MajorityChecker struct {
	dp [][]*params // left-right interval
}

type params struct {
	times int
	value int
}

// Constructor arr[i] >= 1
func Constructor(arr []int) MajorityChecker {
	dp := make([][]*params, len(arr))
	for i := range dp {
		dp[i] = make([]*params, len(arr))
	}

	for left := 0; left < len(arr); left++ {
		dp[left][left] = &params{
			times: 1,
			value: arr[left],
		}

		timesMap := make(map[int]int, len(arr)-left-1)
		timesMap[arr[left]] = 1

		for right := left + 1; right < len(arr); right++ {
			v, _ := timesMap[arr[right]]
			timesMap[arr[right]] = v + 1

			appearTimes := 0
			mostValue := -1
			for key, value := range timesMap {
				if value > appearTimes {
					appearTimes = value
					mostValue = key
				}
			}

			dp[left][right] = &params{
				times: appearTimes,
				value: mostValue,
			}
		}
	}

	return MajorityChecker{
		dp: dp,
	}
}

// Query threshold always bigger than half of left-right interval
// left and right are always valid, ignore params check
func (m *MajorityChecker) Query(left int, right int, threshold int) (res int) {
	param := m.dp[left][right]
	if param.times >= threshold {
		res = param.value
	} else {
		res = -1
	}

	return
}

func contains(intSlice []int, value int) bool {
	containsFlag := false
	for i := range intSlice {
		if intSlice[i] == value {
			containsFlag = true
			break
		}
	}

	return containsFlag
}
