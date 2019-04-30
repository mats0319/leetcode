package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compareOnStringSlice(letterCombinations("23"),
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}))
}

func compareOnStringSlice(a, b []string) bool {
	var (
		equal = true
		ok    bool
		m     = make(map[string]bool)
	)
	for {
		if len(a) != len(b) {
			equal = false
			break
		}

		for i := range a {
			m[a[i]] = true
		}
		for i := range b {
			if _, ok = m[b[i]]; !ok {
				equal = false
				break
			}
		}
		break
	}
	return equal
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var (
		matchArray = [][]string{{""}, {""}, {"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"},
			{"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}
		digit       int
		digitsArray [][]string
		err         error
	)
	if digit, err = strconv.Atoi(digits); err != nil {
		panic(err)
	}
	for ; digit > 0; digit /= 10 {
		digitsArray = append(digitsArray, matchArray[digit%10])
	}

	return recursion(digitsArray)
}

func recursion(ss [][]string) []string {
	var (
		str  []string
		last []string
	)

	str = ss[0]
	if len(ss) == 1 {
		return str
	}

	last = recursion(ss[1:])
	result := make([]string, len(last)*len(str))
	for i := range str {
		for j := range last {
			result[i*len(last)+j] = last[j] + str[i]
		}
	}

	return result
}
