package mario

var keys = []string{"type", "color", "name"} // global variable for leetcode

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    res := 0
    keyIndex := getKeyIndex(ruleKey)
    for i := range items {
        if items[i][keyIndex] == ruleValue {
            res++
        }
    }

    return res
}

func getKeyIndex(key string) int {
    index := -1
    for i := range keys {
        if keys[i] == key {
            index = i
            break
        }
    }

    return index
}
