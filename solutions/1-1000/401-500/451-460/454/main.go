package mario

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sum := make(map[int]int)
	for _, va := range A {
		for _, vb := range B {
			sum[-va-vb]++
		}
	}

	count := 0
	for _, vc := range C {
		for _, vd := range D {
			if v, ok := sum[vc+vd]; ok {
				count += v
			}
		}
	}

	return count
}
