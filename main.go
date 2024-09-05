package main

import "fmt"

// find the greatest common divisor (GCD)
func gcd(a *int, b *int) *int {
	for *b > 0 {
		t := *a
		*a = *b
		*b = t % *b
	}
	return a
}

func main() {
	a := 20
	b := 8
	fmt.Println(*gcd(&a, &b)) // Pass the pointers to 'a' and 'b', dereference the result
}
