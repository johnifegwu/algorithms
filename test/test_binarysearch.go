package test

import (
	"testing"

	"github.com/johnifegwu/algorithms/binarysearch"
)

func TestBinarySearch_ExistingItem(t *testing.T) {
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemToSearch := 5

	index, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err != nil || index != 4 {
		t.Errorf("Expected index 4, but got %d with error: %v", index, err)
	}
}

func TestBinarySearch_NonExistingItem(t *testing.T) {
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemToSearch := 11

	_, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err == nil || err.Error() != "search failed" {
		t.Errorf("Expected error 'search failed', but got %v", err)
	}
}
