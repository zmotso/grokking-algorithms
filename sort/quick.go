package sort

func Quick(toSort []int) {
	if len(toSort) < 2 {
		return
	}

	left, right := 0, len(toSort)-1

	// Pile elements smaller than the pivot on the left
	for i := range toSort {
		if toSort[i] < toSort[right] {
			toSort[i], toSort[left] = toSort[left], toSort[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	toSort[left], toSort[right] = toSort[right], toSort[left]

	// Go down the rabbit hole
	Quick(toSort[:left])
	Quick(toSort[left+1:])

}
