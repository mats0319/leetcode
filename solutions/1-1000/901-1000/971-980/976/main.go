package mario

func largestPerimeter(A []int) int {
    if len(A) < 3 {
        return 0
    }

    QuickSort(A)

    max := 0
    for i := len(A)-1; i > 1; i-- {
        j := i-1
        k := j-1
        if A[j] + A[k] <= A[i] {
            continue
        }

        max = A[i] + A[j] + A[k]
        break
    }

    return max
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
