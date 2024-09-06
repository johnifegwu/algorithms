package stack

// Stack structure using a generic slice
type Stack[T any] struct {
	elements []T
}

// Push adds an element onto the stack
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false // Return zero value of T and false if stack is empty
	}
	// Get the top element
	topIndex := len(s.elements) - 1
	top := s.elements[topIndex]
	// Remove the top element
	s.elements = s.elements[:topIndex]
	return top, true
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false // Return zero value of T and false if stack is empty
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

// Dump the stack
func (s *Stack[T]) Dump() []T {
	return s.elements
}
