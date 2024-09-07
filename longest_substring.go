package main

import "fmt"

func longestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Map to store the last index of each character
	charIndex := make(map[byte]int)
	result := 0
	start := 0

	for end := 0; end < n; end++ {
		// If the character is already in the map and its index is within the current window
		if idx, found := charIndex[s[end]]; found && idx >= start {
			start = idx + 1
		}
		charIndex[s[end]] = end
		// Update the result length
		result = max(result, end-start+1)
	}

	return result
}

func main() {
	rs := longestSubstring("abcadbd")
	fmt.Println(rs)
}
