package mario

var m map[byte]string = map[byte]string{
	'R': "Radiant",
	'D': "Dire",
}

func predictPartyVictory(senate string) string {
	ban := make([]bool, len(senate))
	banR, banD := 0, 0
	for leftR, leftD := true, true; leftR && leftD; {
		for i := 0; i < len(senate); i++ {
			if ban[i%len(senate)] == true {
				continue
			}

			switch senate[i%len(senate)] {
			case 'R':
				if banR > 0 {
					banR--
					ban[i%len(senate)] = true
				} else {
					banD++
				}
			case 'D':
				if banD > 0 {
					banD--
					ban[i%len(senate)] = true
				} else {
					banR++
				}
			}
		}

		leftR, leftD = false, false
		for i := 0; i < len(ban); i++ {
			if !ban[i] {
				if senate[i] == 'R' {
					leftR = true
				} else if senate[i] == 'D' {
					leftD = true
				}

				if leftR && leftD {
					break
				}
			}
		}
	}

	var winner int
	for i, v := range ban {
		if !v {
			winner = i
			break
		}
	}

	return m[senate[winner]]
}
