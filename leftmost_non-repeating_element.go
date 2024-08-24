package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := "geeksforgeeks"
	s2 := "abcd"
	s3 := "abcabc"

	fmt.Printf("The index of the leftmost repeating character in '%s' is: %d\n", s1, leftMostNonRepeatingCharacter(s1))
	fmt.Printf("The index of the leftmost repeating character in '%s' is: %d\n", s2, leftMostNonRepeatingCharacter(s2))
	fmt.Printf("The index of the leftmost repeating character in '%s' is: %d\n", s3, leftMostNonRepeatingCharacter(s3))
}

func leftMostNonRepeatingCharacter(str string) int {
	const CHAR_RANGE = 256
	firstOccurrence := make([]int, CHAR_RANGE)
	for i := 0; i < CHAR_RANGE; i++ {
		firstOccurrence[i] = -1
	}

	for idx := 0; idx < len(str); idx++ {
		if firstOccurrence[str[idx]] == -1 {
			firstOccurrence[str[idx]] = idx
		} else {
			firstOccurrence[str[idx]] = -2
		}
	}

	result := math.MaxInt

	for idx := 0; idx < CHAR_RANGE; idx++ {
		if firstOccurrence[idx] >= 0 {
			result = min(firstOccurrence[idx], result)
		}
	}

	return result
}
