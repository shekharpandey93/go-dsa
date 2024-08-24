package main

import "fmt"

// Function to build the LPS array
func computeLPSArray(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// KMP search function
func KMPSearch(text, pattern string) {
	n := len(text)
	m := len(pattern)
	lps := computeLPSArray(pattern)
	i, j := 0, 0
	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			fmt.Printf("Found pattern at index %d\n", i-j)
			j = lps[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
}

func main() {
	text := "ababcababcabc"
	pattern := "ababc"
	KMPSearch(text, pattern)
}
