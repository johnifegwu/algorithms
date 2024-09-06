package quicksort

import "golang.org/x/exp/constraints"

// QuickSort sorts a generic slice in place using the QuickSort algorithm
func QuickSort[T constraints.Ordered](dataset []T) {
	if len(dataset) <= 1 {
		return
	}

	// Perform the partitioning
	pivotIndex := partition(dataset)

	// Recursively sort the left and right subarrays
	QuickSort(dataset[:pivotIndex])
	QuickSort(dataset[pivotIndex+1:])
}

// partition partitions the slice around a pivot and returns the index of the pivot
func partition[T constraints.Ordered](dataset []T) int {
	pivot := dataset[len(dataset)-1] // Use the last element as the pivot
	i := 0                           // Index of the smaller element

	for j := 0; j < len(dataset)-1; j++ {
		if dataset[j] < pivot {
			dataset[i], dataset[j] = dataset[j], dataset[i]
			i++
		}
	}

	// Swap the pivot element to its correct position
	dataset[i], dataset[len(dataset)-1] = dataset[len(dataset)-1], dataset[i]

	return i
}
