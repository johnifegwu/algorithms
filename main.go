package main

import (
	"fmt"

	"github.com/johnifegwu/algorithms/algorithms/gcd"
	"github.com/johnifegwu/algorithms/datastructures/hashmap"
	"github.com/johnifegwu/algorithms/datastructures/queue"
	"github.com/johnifegwu/algorithms/datastructures/stack"
)

func main() {

	// HashMap datastructure usage example
	//==================================================
	// Create a new HashMap for strings to integers
	hashMap := hashmap.NewHashMap[string, int]()
	hashMap.Add("one", 1)
	hashMap.Add("two", 2)
	// dump hashmap
	fmt.Println(hashMap.Dump())

	// Remove an existing key
	removed := hashMap.Remove("one")
	fmt.Println("Removed 'one':", removed) // Output: Removed 'one': true

	// Attempt to remove a non-existing key
	removed = hashMap.Remove("three")
	fmt.Println("Removed 'three':", removed) // Output: Removed 'three': false
	// dump hashmap
	fmt.Println(hashMap.Dump())

	// Check the value of a remaining key
	if value, exists := hashMap.Get("two"); exists {
		fmt.Println("Value for 'two':", value) // Output: Value for 'two': 2
	} else {
		fmt.Println("'two' not found")
	}
	//==================================================

	// Queue datastructure usage example
	//==================================================
	// Create a new queue of integers
	intQueue := &queue.Queue[int]{}
	intQueue.Enqueue(10)
	intQueue.Enqueue(20)
	intQueue.Enqueue(30)
	// print queue
	fmt.Println(intQueue.Dump())

	fmt.Println("Queue Size:", intQueue.Size()) // Output: Queue Size: 3

	if front, ok := intQueue.Peek(); ok {
		fmt.Println("Front Element:", front) // Output: Front Element: 10
	}

	if dequeued, ok := intQueue.Dequeue(); ok {
		fmt.Println("Dequeued:", dequeued) // Output: Dequeued: 10
		// print queue
		fmt.Println(intQueue.Dump())
	}

	fmt.Println("Queue Size after Dequeue:", intQueue.Size()) // Output: Queue Size: 2
	//==================================================

	// Stack datastructure usage example
	//==================================================
	// Initialize a new stack for integers
	intStack := &stack.Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	// print stack
	fmt.Println(intStack.Dump())

	// Peek the top element and check if it's valid
	if top, ok := intStack.Peek(); ok {
		fmt.Println("Int Stack Peek:", top) // Output: Int Stack Peek: 30
	} else {
		fmt.Println("Stack is empty.")
	}

	// Pop the top element and check if it's valid
	if popped, ok := intStack.Pop(); ok {
		fmt.Println("Int Stack Pop:", popped) // Output: Int Stack Pop: 30
		// print stack
		fmt.Println(intStack.Dump())
	} else {
		fmt.Println("Stack is empty.")
	}

	// Initialize a new stack for strings
	stringStack := &stack.Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")

	// print stack
	fmt.Println(stringStack.Dump())

	// Peek the top element and check if it's valid
	if top, ok := stringStack.Peek(); ok {
		fmt.Println("String Stack Peek:", top) // Output: String Stack Peek: World
	} else {
		fmt.Println("Stack is empty.")
	}

	// Pop the top element and check if it's valid
	if popped, ok := stringStack.Pop(); ok {
		fmt.Println("String Stack Pop:", popped) // Output: String Stack Pop: World
		// print stack
		fmt.Println(stringStack.Dump())
	} else {
		fmt.Println("Stack is empty.")
	}
	//============================================================

	//Greatest common diniminator implementation
	a := 20
	b := 8
	fmt.Println(*gcd.GetGcd(&a, &b)) // Pass the pointers to 'a' and 'b', dereference the result
}
