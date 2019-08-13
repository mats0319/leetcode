package pn

import "testing"

var testCase = []int{
	121,
	-121,
	10,
}

var result = []bool{
	true,
	false,
	false,
}

func TestIsPalindrome(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if isPalindrome(tcs[i]) != result[i] {
			t.Errorf("palindrome number test failed on case: %d", i)
		}
	}
}
