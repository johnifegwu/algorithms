package test

import (
	"reflect"
	"testing"

	"github.com/johnifegwu/algorithms/datastructures/set"
)

// Test Add, Remove, Contains, Size, Clear, and ToSlice functions for a generic Set

// Test case for adding elements to the Set
func TestSet_Add(t *testing.T) {
	set := set.NewSet[int]()
	set.Add(1)
	set.Add(2)

	if set.Size() != 2 {
		t.Errorf("Expected size 2, got %d", set.Size())
	}

	if !set.Contains(1) || !set.Contains(2) {
		t.Errorf("Expected set to contain 1 and 2")
	}
}

// Test case for removing elements from the Set
func TestSet_Remove(t *testing.T) {
	set := set.NewSet[int]()
	set.Add(1)
	set.Add(2)

	set.Remove(1)

	if set.Contains(1) {
		t.Errorf("Expected set to not contain 1 after removal")
	}

	if set.Size() != 1 {
		t.Errorf("Expected size 1, got %d", set.Size())
	}
}

// Test case for checking if an element is in the Set
func TestSet_Contains(t *testing.T) {
	set := set.NewSet[string]()
	set.Add("apple")
	set.Add("banana")

	if !set.Contains("apple") {
		t.Errorf("Expected set to contain 'apple'")
	}

	if set.Contains("cherry") {
		t.Errorf("Expected set to not contain 'cherry'")
	}
}

// Test case for getting the size of the Set
func TestSet_Size(t *testing.T) {
	set := set.NewSet[float64]()
	set.Add(1.1)
	set.Add(2.2)
	set.Add(3.3)

	if set.Size() != 3 {
		t.Errorf("Expected size 3, got %d", set.Size())
	}
}

// Test case for clearing all elements in the Set
func TestSet_Clear(t *testing.T) {
	set := set.NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Clear()

	if set.Size() != 0 {
		t.Errorf("Expected size 0 after clearing, got %d", set.Size())
	}

	if set.Contains(1) {
		t.Errorf("Expected set to not contain any elements after clearing")
	}
}

// Test case for converting the Set to a slice
func TestSet_ToSlice(t *testing.T) {
	set := set.NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	slice := set.ToSlice()

	// Sort the slice for comparison (optional)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Expected slice %v, got %v", expected, slice)
	}
}

// Test case for handling duplicates in the Set
func TestSet_Duplicates(t *testing.T) {
	set := set.NewSet[int]()
	set.Add(1)
	set.Add(1)
	set.Add(2)

	if set.Size() != 2 {
		t.Errorf("Expected size 2 with no duplicates, got %d", set.Size())
	}

	if !set.Contains(1) || !set.Contains(2) {
		t.Errorf("Expected set to contain 1 and 2")
	}
}

// Test case for handling empty Set
func TestSet_Empty(t *testing.T) {
	set := set.NewSet[int]()

	if set.Size() != 0 {
		t.Errorf("Expected size 0 for an empty set, got %d", set.Size())
	}

	if set.Contains(1) {
		t.Errorf("Expected set to not contain any elements")
	}
}
