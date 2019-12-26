package sort

func Selection(toSort []int) {
	for i := 0; i < len(toSort); i++ {
		var min = i

		for j := i; j < len(toSort); j++ {
			if toSort[j] < toSort[min] {
				min = j
			}
		}

		toSort[i], toSort[min] = toSort[min], toSort[i]
	}
}

func findMin(in []int) int {
	minInd := 0
	minVal := in[minInd]
	for k, v := range in {
		if v < minVal {
			minInd = k
			minVal = v
		}
	}

	return minInd
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]

	return s[:len(s)-1]
}
