package sort

func BubbleSort(is []int) {
    var (
        ordered bool
        length  = len(is)
    )
    if length <= 1 {
        return
    }

    for i := 0; i < length && !ordered; i++ {
        ordered = true
        for j := 0; j < length-i-1; j++ {
            if is[j] > is[j+1] {
                ordered = false
                is[j], is[j+1] = is[j+1], is[j]
            }
        }
    }

    return
}

func SelectionSort(is []int) {
    var (
        ordered  bool
        length   = len(is)
        maxIndex int
    )
    if length <= 1 {
        return
    }

    for i := 0; i < length && !ordered; i++ {
        ordered = true
        maxIndex = 0
        for j := 0; j < length-i; j++ {
            if is[j] >= is[maxIndex] {
                ordered = false
                maxIndex = j
            }
        }
        if maxIndex != length-i-1 {
            is[maxIndex], is[length-i-1] = is[length-i-1], is[maxIndex]
        }
    }

    return
}

func QuickSort(is []int) {
    if len(is) <= 1 {
        return
    }

    var small int
    {
        var big int
        for i := 1; i < len(is); i++ {
            if is[i] > is[0] {
                big++
            } else {
                small++
                if big != 0 {
                    is[i], is[small] = is[small], is[i]
                }
            }
        }
        if small != 0 {
            is[0], is[small] = is[small], is[0]
        }
    }

    QuickSort(is[:small])
    QuickSort(is[small+1:])

    return
}

func CountingSort(is []int) {
    var (
        countingArray []int
        min int
    )
    {
        var max, length = 0, len(is)
        {
            if length <= 1 {
                return
            }

            max, min = is[0], is[0]
            for i := 1; i < length; i++ {
                if is[i] > max {
                    max = is[i]
                } else if is[i] < min {
                    min = is[i]
                }
            }
        }

        countingArray = make([]int, max-min+1)
        for i := 0; i < length; i++ {
            countingArray[is[i]-min]++
        }
    }

    var index int
    for i := 0; i < len(countingArray); i++ {
        for ; countingArray[i] > 0; countingArray[i]-- {
            is[index] = i + min
            index++
        }
    }

    return
}

func InsertionSort(is []int) {
    length := len(is)
    if length <= 1 {
        return
    }

    var pos, temp int
    for i := 1; i < length; i++ {
        if is[i] >= is[i-1] {
            continue
        }
        for pos = i; pos > 0 && is[i] < is[pos-1]; pos-- {
        }
        temp = is[i]
        for t := i; t > pos; t-- {
            is[t] = is[t-1]
        }
        is[pos] = temp
    }

    return
}
