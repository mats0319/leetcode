package mario

func mostVisited(n int, rounds []int) []int {
    isOneMoreVisited := false
    for i := 0; i < len(rounds) - 1; i++ {
        if rounds[i] > rounds[i+1] {
            isOneMoreVisited = true
            break
        }
    }

    res := make([]int, 0, n)
    start, end := rounds[0], rounds[len(rounds)-1]
    if !isOneMoreVisited || isOneMoreVisited && start <= end { // less than one circle
        for i := start; i <= end; i++ {
            res = append(res, i)
        }
    } else if start > end {
        for i := 1; i <= n; i++ {
           if i > end && i < start {
               continue
           }

           res = append(res, i)
        }
    }

    return res
}
