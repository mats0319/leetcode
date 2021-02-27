package mario

// customers.length >= 1
func averageWaitingTime(customers [][]int) float64 {
	waitTime := 0
	{
        currTime := 0
        for i := 0; i < len(customers); i++ {
            if currTime < customers[i][0] {
                currTime = customers[i][0]
            }

            currTime += customers[i][1]
            waitTime += currTime-customers[i][0]
        }
    }

	return float64(waitTime) / float64(len(customers))
}
