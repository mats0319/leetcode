package mario

func partitionLabels(S string) []int {
    last := makeArrays(S)

    res := make([]int, 0, 26) // res.length will <= 26
    for i := 0; i < len(S); {
        end := last[S[i]-'a']
        for start := i+1; start < end; start++ {
            if last[S[start]-'a'] > end {
                end = last[S[start]-'a']
            }
        }
        res = append(res, end-i+1)
        i = end+1
    }

    return res
}

func makeArrays(str string) [26]int {
    start := [26]int{}
    end := [26]int{}
    for i := 0; i < len(start); i++ {
        start[i] = -1
        end[i] = -1
    }

    for i := range str {
        last := len(str)-1
        index := str[i]-'a'

        if start[index] != -1 {
            continue
        }

        start[index] = i
        for i <= last {
            if str[last] == str[i] {
                end[index] = last
                break
            }

            last--
        }
    }

    return end
}
