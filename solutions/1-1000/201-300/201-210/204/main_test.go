package mario

import "testing"

var testCase = []struct {
	In     int
	Expect int
}{
	// test cases here
	{10, 4},
	{3, 1},
}

func TestCountPrimes(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if countPrimes(tcs[i].In) != tcs[i].Expect {
			t.Errorf("count primes test failed on case: %d\n", i)
		}
	}
}
