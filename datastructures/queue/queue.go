package queue

// Queue structure using a generic slice
type Queue[T any] struct {
	elements []T
}

// Enqueue adds an element at the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the first element from the queue
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, false // Return zero value of T and false if queue is empty
	}
	// Get the first element
	first := q.elements[0]
	// Remove the first element (efficient slicing)
	q.elements = q.elements[1:]
	return first, true
}

// Peek returns the first element without removing it
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, false // Return zero value of T and false if queue is empty
	}
	return q.elements[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return len(q.elements)
}

// Dump Queue
func (q *Queue[T]) Dump() []T {
	return q.elements
}
