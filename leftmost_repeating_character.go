package main

import "fmt"

func leftMostRepeatingCharacter(s string) int {
	const CHAR_RANGE = 256
	// Initialize an array to track whether a character has been seen
	seen := make([]bool, CHAR_RANGE)

	// Initialize the result as -1
	result := -1

	// Traverse the string from right to left (reverse order)
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if seen[ch] {
			// If the character is already seen, update the result to current index
			result = i
		} else {
			// Mark the character as seen
			seen[ch] = true
		}
	}

	return result
}

func main() {
	s1 := "geeksforgeeks"
	s2 := "abcd"
	s3 := "abcabc"

	fmt.Printf("The index of the leftmost repeating character in '%s' is: %d\n", s1, leftMostRepeatingCharacter(s1))
	fmt.Printf("The index of the leftmost repeating character in '%s' is: %d\n", s2, leftMostRepeatingCharacter(s2))
	fmt.Printf("The index of the leftmost repeating character in '%s' is: %d\n", s3, leftMostRepeatingCharacter(s3))
}
