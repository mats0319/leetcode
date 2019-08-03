package lswrc

import (
    "testing"
)

var testCase = []string{
    "abcabcbba",
    "bbbbb",
    "pwwkew",
    " ",
    "aab",
}

var result = []int{
    3,
    1,
    3,
    1,
    2,
}

func TestLengthOfLongestSubstring(t *testing.T) {
    tcs := testCase
    for i := range tcs {
        if lengthOfLongestSubstring(tcs[i]) != result[i] {
            t.Errorf("length of longest substring test failed on case: %d", i)
        }
    }
}
