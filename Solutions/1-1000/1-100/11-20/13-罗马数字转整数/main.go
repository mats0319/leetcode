package main

// map init
var m = map[string]int{
	"I":  1,
	"IV": 3,
	"V":  5,
	"IX": 8,
	"X":  10,
	"XL": 30,
	"L":  50,
	"XC": 80,
	"C":  100,
	"CD": 300,
	"D":  500,
	"CM": 800,
	"M":  1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	result := m[s[0:1]]
	for i := 1; i < len(s); i++ {
		if v, ok := m[s[i-1:i+1]]; ok {
			result += v
		} else {
			result += m[s[i:i+1]]
		}
	}

	return result
}
