package mario

// cost.length >= 2
func minCostClimbingStairs(cost []int) int {
	dp := []int{cost[0], cost[1], 0}
	for i := 2; i < len(cost); i++ {
		dp[i%3] = min(dp[(i-1)%3], dp[(i-2)%3]) + cost[i]
	}

	return min(dp[(len(cost)+1)%3], dp[(len(cost)+2)%3])
}

func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}

	return
}
