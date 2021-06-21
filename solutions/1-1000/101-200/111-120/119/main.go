package mario

func getRow(rowIndex int) []int {
    line := []int{1}

    for i := 0; i < rowIndex; i++ {
        nextLine := make([]int, len(line)+1)
        nextLine[0] = 1
        nextLine[len(nextLine)-1] = 1

        for j := 1; j < len(nextLine)-1; j++ {
            nextLine[j] = line[j]+line[j-1]
        }

        line = nextLine
    }

    return line
}
