package test

import (
	"testing"

	"github.com/johnifegwu/algorithms/algorithms/recursion"
)

// TestFindMaxValue_SingleElement checks the behavior for a single-element slice
func TestFindMaxValue_SingleElement(t *testing.T) {
	dataset := []int{42}
	result := recursion.FindMaxValue(dataset)
	expected := 42

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestFindMaxValue_MultipleElements checks the behavior for a general slice
func TestFindMaxValue_MultipleElements(t *testing.T) {
	dataset := []int{1, 3, 5, 7, 9, 2, 4, 6, 8}
	result := recursion.FindMaxValue(dataset)
	expected := 9

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestFindMaxValue_NegativeNumbers checks the behavior when the slice contains negative numbers
func TestFindMaxValue_NegativeNumbers(t *testing.T) {
	dataset := []int{-10, -3, -50, -7, -1}
	result := recursion.FindMaxValue(dataset)
	expected := -1

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestFindMaxValue_AllEqualElements checks the behavior when all elements in the slice are the same
func TestFindMaxValue_AllEqualElements(t *testing.T) {
	dataset := []int{7, 7, 7, 7, 7}
	result := recursion.FindMaxValue(dataset)
	expected := 7

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestFindMaxValue_EmptySlice checks the behavior when the slice is empty
func TestFindMaxValue_EmptySlice(t *testing.T) {
	dataset := []int{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty slice but function did not panic")
		}
	}()

	recursion.FindMaxValue(dataset)
}

// TestFindMaxValue_LargeDataset checks the behavior for a large dataset
func TestFindMaxValue_LargeDataset(t *testing.T) {
	dataset := []int{1, 1000, 50, 100, 200, 500, 10000, 3000}
	result := recursion.FindMaxValue(dataset)
	expected := 10000

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
