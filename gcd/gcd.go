package gcd

// find the greatest common divisor (GCD)
func GetGcd(a *int, b *int) *int {
	for *b > 0 {
		t := *a
		*a = *b
		*b = t % *b
	}
	return a
}
