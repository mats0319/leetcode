package mario

func jump(nums []int) int {
    return recurse(nums, 0)
}

// recurse input left array and steps used
func recurse(array []int, steps int) int {
    if len(array) == 1 {
        return 0
    }

    steps++

    res := -1
    for i := array[0]; i >= 1; i-- {
        if i >= len(array)-1 { // can jump to the end
            res = steps
            break
        }

        minSteps := recurse(array[i:], steps)
        if res < 0 || (0 <= minSteps && minSteps < res) {
            res = minSteps
        }
    }

    return res
}
