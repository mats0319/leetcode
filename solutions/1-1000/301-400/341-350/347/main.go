package mario

type Node struct {
    Value int
    Frequency int

    Next *Node
}

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int) // map, element - frequency
    for i := range nums {
        m[nums[i]]++
    }

    // rule: value from small to big, find it's position and modify in-place
    resList := &Node{}
    result := make([]int, k) // define before use k
    for value, frequency := range m {
        k--
        resList.Value = value
        resList.Frequency = frequency
        delete(m, value)

        break // run only once
    }
    for value, frequency := range m {
        if k > 0 {
            resList = resList.Upsert(value, frequency, false)

            k--
            delete(m, value)
        }
    }
    for value, frequency := range m {
        if frequency > resList.Frequency {
            resList = resList.Upsert(value, frequency, true)
        }
    }

    for i := range result {
        result[i] = resList.Value
        resList = resList.Next
    }

    return result
}

// list is ordered from small to big, sort by frequency
func (n *Node) Upsert(value, frequency int, update bool) (result *Node) {
    pre, curr := &Node{Next: n}, n
    result = pre
    for ; curr != nil && frequency > curr.Frequency; pre, curr = pre.Next, curr.Next {
    }

    // curr == nil is ok
    pre.Next = &Node{Value: value, Frequency: frequency, Next: curr}

    result = result.Next
    if update {
        result = result.Next
    }

    return
}
