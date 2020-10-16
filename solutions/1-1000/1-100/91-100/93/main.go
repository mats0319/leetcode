package mario

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 16 {
		return nil
	}

	res := make([]string, 0)
	// '[start, end)'
	for a := 1; a < 4; a++ {
		for b := a+1; b < len(s)-1 && b < a+4; b++ {
			for c := b+1; c < len(s) && c < b+4; c++ {
				if len(s[c:]) > 3 {
					continue
				}

				str := fmt.Sprintf("%s.%s.%s.%s", s[0:a], s[a:b], s[b:c], s[c:])
				if isValid(str) {
					res = append(res, str)
				}
			}
		}
	}

	return res
}

// make sure ip has 4 part
func isValid(ip string) (isValidFlag bool) {
	split := strings.Split(ip, ".")

	isValidFlag = true
	for i := 0; i < len(split); i++ {
		value, err := strconv.Atoi(split[i])
		if has0Prefix(split[i]) || err != nil || !validValue(value) {
			isValidFlag = false
			break
		}
	}

	return
}

func has0Prefix(s string) bool {
	return len(s) > 1 && s[0] == '0'
}

func validValue(v int) bool {
	return 0 <= v && v <= 255
}
