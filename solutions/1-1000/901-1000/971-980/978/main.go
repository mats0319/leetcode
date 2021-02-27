package mario

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}

	maxLength := 1
	{
		lastTrend := arr[1] - arr[0]
		var length int
		if lastTrend == 0 {
		    length = 1
        } else {
            length = 2
        }

		for i := 2; i < len(arr); i++ {
			currTrend := arr[i] - arr[i-1]

			if isOpposite(lastTrend, currTrend) {
			    lastTrend = currTrend
			    length++
			    continue
            }

            if length > maxLength {
                maxLength = length
            }

            if currTrend == 0 {
                length = 1
            } else {
                length = 2 // 这里不兼容初始化lt = 0和后续lt = 0的情况，所以整体从i = 2开始处理
            }

			lastTrend = currTrend
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func isOpposite(a, b int) bool {
	return (a > 0 && b < 0) || (a < 0 && b > 0)
}
