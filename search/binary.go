package search

import "errors"

var ErrNotFount = errors.New("not found in slice")

func Binary(search int, in []int) (int, error) {
	low := 0
	high := len(in) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := in[mid]
		if guess == search {
			return mid, nil
		}
		if guess < search {
			low += mid
		} else {
			high -= mid
		}
	}

	return 0, ErrNotFount
}
