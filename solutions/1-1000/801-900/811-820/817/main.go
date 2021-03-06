package mario

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	gMap := make(map[int]bool, len(G))
	for i := range G {
		gMap[G[i]] = true
	}

	flag := make([]bool, 0) // attention: define on length
	for p := head; p != nil; p = p.Next {
		flag = append(flag, gMap[p.Val])
	}

	maxLength := 0
	containsFlag := false
	count := 0
	for i := range flag {
		containsFlag = flag[i]
		if containsFlag {
			count++
		}

		if count > maxLength {
			maxLength = count
		}

		if !containsFlag {
			count = 0
		}
	}

	return maxLength
}
