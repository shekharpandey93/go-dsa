package main

import (
	"fmt"
)

func numTrailingZeros(n int) int {

	// This function return the number of
	// trailing zeros in the factorial of
	// a number n.
	// The count variable denotes the
	// number of trailing zeros in n!
	if n < 0 {
		return -1
	}
	count := 0
	for i := 5; n/i >= 1; i *= 5 {
		count += n / i
	}
	return count
}

func main() {

	//Testcases:
	// Testcase 1: n:= 10
	// Testcase 2: n:= 100
	// Testcase 3: n:= 125
	// Testcase 4: n:= 530

	n := 100
	result := numTrailingZeros(n)
	fmt.Printf("The number of trailing zeros"+" in %v! is %v", n, result)
}
