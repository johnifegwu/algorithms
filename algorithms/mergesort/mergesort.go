package mergesort

import "golang.org/x/exp/constraints"

// MergeSort performs a merge sort on a generic slice
func MergeSort[T constraints.Ordered](dataset []T) []T {
	if len(dataset) > 1 {
		// Divide dataset into two halves
		mid := len(dataset) / 2
		leftdata := MergeSort(dataset[:mid])
		rightdata := MergeSort(dataset[mid:])

		// Perform the merging
		return merge(leftdata, rightdata)
	}
	return dataset
}

// merge merges two sorted slices into a single sorted slice
func merge[T constraints.Ordered](left, right []T) []T {
	i, j := 0, 0
	result := make([]T, 0, len(left)+len(right))

	// While both datasets have content
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining values from the left dataset
	result = append(result, left[i:]...)

	// Append any remaining values from the right dataset
	result = append(result, right[j:]...)

	return result
}
