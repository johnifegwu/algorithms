package main

import (
	"fmt"

	"github.com/johnifegwu/algorithms/datastructures/stack"
	"github.com/johnifegwu/algorithms/gcd"
)

func main() {

	// Stack datastructure usage example
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
