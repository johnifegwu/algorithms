package hashmap

type HashMap[Key string, Value any] struct {
	elements map[Key]Value
}

// NewHashMap creates a new HashMap instance
func NewHashMap[Key string, Value any]() *HashMap[Key, Value] {
	return &HashMap[Key, Value]{
		elements: make(map[Key]Value),
	}
}

// Add inserts a key-value pair into the HashMap
func (h *HashMap[Key, Value]) Add(key Key, value Value) {
	h.elements[key] = value
}

// Remove deletes a key-value pair from the HashMap
func (h *HashMap[Key, Value]) Remove(key Key) bool {
	if _, exists := h.elements[key]; exists {
		delete(h.elements, key)
		return true // Successfully removed the element
	}
	return false // Key does not exist
}

// Get retrieves a value by its key
func (h *HashMap[Key, Value]) Get(key Key) (Value, bool) {
	value, exists := h.elements[key]
	return value, exists
}

// dump hashmap
func (h *HashMap[Key, Value]) Dump() []Value {
	values := make([]Value, len(h.elements))
	i := 0
	for _, v := range h.elements {
		values[i] = v
		i++
	}

	return values
}
