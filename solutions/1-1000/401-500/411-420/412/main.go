package mario

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		str := strconv.Itoa(i)
		{
			t := ""
			if i%3 == 0 {
				t += "Fizz"
			}
			if i%5 == 0 {
				t += "Buzz"
			}

			if t != "" {
				str = t
			}
		}

		result[i-1] = str
	}

	return result
}
