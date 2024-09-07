package binarysearch

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](dataset []T, itemToSearch T) (index int, err error) {
	datasetLen := len(dataset)
	lowerIndex := 0
	uperIndex := datasetLen - 1

	for lowerIndex <= uperIndex {

		// get the midpoint
		midpoint := (lowerIndex + uperIndex) / 2

		// return the index if item is found
		if dataset[midpoint] == itemToSearch {
			return midpoint, nil
		}

		// otherwise advance the midpoint
		if itemToSearch > dataset[midpoint] {
			lowerIndex = midpoint + 1
		} else {
			uperIndex = midpoint - 1
		}
	}

	return 0, errors.New("search failed")
}
