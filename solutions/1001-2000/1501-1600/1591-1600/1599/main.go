package mario

// customers.length >= 1
func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	if boardingCost*4 <= runningCost {
		return -1
	}

	customersLeft := 0
	profit := 0
	maxProfit := 0
	result := -1
	for i := 0; customersLeft > 0 || i < len(customers); i++ {
		if i < len(customers) {
			customersLeft += customers[i]
		}

		up := 0
		{
			if customersLeft > 4 {
				up = 4
				customersLeft -= 4
			} else {
				up = customersLeft
				customersLeft = 0
			}
		}

		profit += up*boardingCost - runningCost
		if profit > maxProfit { // min operation, not equal, bigger only
			maxProfit = profit
			result = i + 1 // index to times
		}
	}

	return result
}
