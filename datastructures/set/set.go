package set

import (
	"golang.org/x/exp/constraints"
)

type Set[T constraints.Ordered] struct {
	elements map[T]struct{}
}

func NewSet[T constraints.Ordered]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// checks if the provided value exists in the set
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// adds an element to the set
func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// removes an element from the set
func (s *Set[T]) Remove(element T) bool {
	if _, exist := s.elements[element]; exist {
		delete(s.elements, element)
		return true // exist and removed
	}
	return false // does not exist
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Clear removes all elements from the set
func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// ToSlice converts the set to a slice of elements
func (s *Set[T]) ToSlice() []T {
	Keys := make([]T, 0, len(s.elements))
	for k := range s.elements {
		Keys = append(Keys, k)
	}
	return Keys
}
