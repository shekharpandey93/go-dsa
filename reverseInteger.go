package main

import "fmt"

func main() {
	fmt.Println(Reverse(12345))  // Output: 54321
	fmt.Println(Reverse(-12345)) // Output: -54321
	fmt.Println(Reverse(120))    // Output: 21
	fmt.Println(Reverse(0))      // Output: 0
}

func Reverse(x int) int {
	isNegative := false
	result := 0
	if x < 0 {
		isNegative = true
		x = -x
	}
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}

	if isNegative {
		result *= -1
	}

	return result
}
