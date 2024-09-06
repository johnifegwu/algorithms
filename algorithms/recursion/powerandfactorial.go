package recursion

// Power calculates num^pwr recursively
// 2^4 = 2*2*2*2 = 16
func Power(num int, pwr int) int {
	if pwr == 0 {
		return 1
	} else {
		return (num * Power(num, pwr-1))
	}
}

// Factorial calculates num! recursively
// 5! = 5*4*6*2*1
// Special case 0! is 1, because ... math
func Factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return (num * Factorial(num-1))
	}
}
