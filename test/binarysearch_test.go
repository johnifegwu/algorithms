package test

import (
	"testing"

	"github.com/johnifegwu/algorithms/algorithms/binarysearch"
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

func TestBinarySearch_EmptyDataset(t *testing.T) {
	dataset := []int{}
	itemToSearch := 1

	_, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err == nil || err.Error() != "search failed" {
		t.Errorf("Expected error 'search failed', but got %v", err)
	}
}

func TestBinarySearch_FirstItem(t *testing.T) {
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemToSearch := 1

	index, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err != nil || index != 0 {
		t.Errorf("Expected index 0, but got %d with error: %v", index, err)
	}
}

func TestBinarySearch_LastItem(t *testing.T) {
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemToSearch := 10

	index, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err != nil || index != 9 {
		t.Errorf("Expected index 9, but got %d with error: %v", index, err)
	}
}

func TestBinarySearch_UnsortedDataset(t *testing.T) {
	dataset := []int{3, 1, 4, 2}
	itemToSearch := 3

	_, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err == nil || err.Error() != "search failed" {
		t.Errorf("Expected error 'search failed' due to unsorted dataset, but got %v", err)
	}
}

func TestBinarySearch_StringDataset(t *testing.T) {
	dataset := []string{"apple", "banana", "cherry", "date", "fig"}
	itemToSearch := "cherry"

	index, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err != nil || index != 2 {
		t.Errorf("Expected index 2, but got %d with error: %v", index, err)
	}
}

func TestBinarySearch_Duplicates(t *testing.T) {
	dataset := []int{1, 2, 2, 2, 3}
	itemToSearch := 2

	index, err := binarysearch.BinarySearch(dataset, itemToSearch)

	if err != nil || (index < 1 || index > 3) {
		t.Errorf("Expected index between 1 and 3, but got %d with error: %v", index, err)
	}
}

func TestIsSortedAsc_EmptyDataset(t *testing.T) {
	dataset := []int{}
	if !binarysearch.IsSortedAsc(dataset) {
		t.Errorf("Expected true for empty dataset")
	}
}

func TestIsSortedAsc_SortedDataset(t *testing.T) {
	dataset := []int{1, 2, 3, 4, 5}
	if !binarysearch.IsSortedAsc(dataset) {
		t.Errorf("Expected true for sorted dataset")
	}
}

func TestIsSortedAsc_UnsortedDataset(t *testing.T) {
	dataset := []int{5, 3, 2, 1}
	if binarysearch.IsSortedAsc(dataset) {
		t.Errorf("Expected false for unsorted dataset")
	}
}

func TestIsSortedAsc_SingleElementDataset(t *testing.T) {
	dataset := []int{42}
	if !binarysearch.IsSortedAsc(dataset) {
		t.Errorf("Expected true for single element dataset")
	}
}

func TestIsSortedAsc_DuplicateValues(t *testing.T) {
	dataset := []int{1, 2, 2, 3, 4}
	if !binarysearch.IsSortedAsc(dataset) {
		t.Errorf("Expected true for dataset with duplicates")
	}
}

func TestIsSortedAsc_FloatDataset(t *testing.T) {
	dataset := []float64{1.1, 2.2, 3.3, 4.4}
	if !binarysearch.IsSortedAsc(dataset) {
		t.Errorf("Expected true for sorted float dataset")
	}
}

func TestIsSortedDesc_EmptyDataset(t *testing.T) {
	dataset := []int{}
	if !binarysearch.IsSortedDesc(dataset) {
		t.Errorf("Expected true for empty dataset")
	}
}

func TestIsSortedDesc_SortedDataset(t *testing.T) {
	dataset := []int{5, 4, 3, 2, 1}
	if !binarysearch.IsSortedDesc(dataset) {
		t.Errorf("Expected true for sorted dataset")
	}
}

func TestIsSortedDesc_UnsortedDataset(t *testing.T) {
	dataset := []int{1, 3, 2, 4}
	if binarysearch.IsSortedDesc(dataset) {
		t.Errorf("Expected false for unsorted dataset")
	}
}

func TestIsSortedDesc_SingleElementDataset(t *testing.T) {
	dataset := []int{42}
	if !binarysearch.IsSortedDesc(dataset) {
		t.Errorf("Expected true for single element dataset")
	}
}

func TestIsSortedDesc_DuplicateValues(t *testing.T) {
	dataset := []int{5, 5, 4, 3, 2}
	if !binarysearch.IsSortedDesc(dataset) {
		t.Errorf("Expected true for dataset with duplicates")
	}
}

func TestIsSortedDesc_FloatDataset(t *testing.T) {
	dataset := []float64{4.4, 3.3, 2.2, 1.1}
	if !binarysearch.IsSortedDesc(dataset) {
		t.Errorf("Expected true for sorted float dataset")
	}
}
