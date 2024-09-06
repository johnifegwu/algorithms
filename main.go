package main

import (
	"fmt"

	"github.com/johnifegwu/algorithms/algorithms/gcd"
	"github.com/johnifegwu/algorithms/algorithms/mergesort"
	"github.com/johnifegwu/algorithms/algorithms/recursion"
	"github.com/johnifegwu/algorithms/datastructures/hashmap"
	"github.com/johnifegwu/algorithms/datastructures/queue"
	"github.com/johnifegwu/algorithms/datastructures/stack"
)

func main() {
	// Algorithm test
	//=============================================================

	// Merge Sort
	// Test MergeSort with integers
	data := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted int slice:", data)
	sortedData := mergesort.MergeSort(data)
	fmt.Println("Sorted int slice:", sortedData)

	// Test MergeSort with floats
	dataFloat := []float64{3.14, 2.71, 1.41, 1.73}
	fmt.Println("\nUnsorted float slice:", dataFloat)
	sortedDataFloat := mergesort.MergeSort(dataFloat)
	fmt.Println("Sorted float slice:", sortedDataFloat)

	// Test MergeSort with strings
	dataString := []string{"banana", "apple", "cherry", "date"}
	fmt.Println("\nUnsorted string slice:", dataString)
	sortedDataString := mergesort.MergeSort(dataString)
	fmt.Println("Sorted string slice:", sortedDataString)

	// Recursion
	// Test Power function
	fmt.Println("Testing Power Function:")
	base := 2
	exponent := 4
	result := recursion.Power(base, exponent)
	fmt.Printf("%d^%d = %d\n", base, exponent, result) // Output: 2^4 = 16

	// Test Power with different inputs
	base = 5
	exponent = 3
	result = recursion.Power(base, exponent)
	fmt.Printf("%d^%d = %d\n", base, exponent, result) // Output: 5^3 = 125

	// Test Factorial function
	fmt.Println("\nTesting Factorial Function:")
	num := 5
	factResult := recursion.Factorial(num)
	fmt.Printf("%d! = %d\n", num, factResult) // Output: 5! = 120

	// Test Factorial with 0
	num = 0
	factResult = recursion.Factorial(num)
	fmt.Printf("%d! = %d\n", num, factResult) // Output: 0! = 1

	//Greatest common diniminator implementation
	a := 20
	b := 8
	fmt.Println(*gcd.GetGcd(&a, &b)) // Pass the pointers to 'a' and 'b', dereference the result

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
}
