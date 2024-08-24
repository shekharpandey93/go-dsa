package main

import "fmt"

func main() {
	s := "eceba"
	k := 2
	fmt.Printf("The length of the longest substring with %d distinct characters is: %d\n", k, lengthOfLongestSubstring(s, k))
	fmt.Printf("The length of the longest substring with %d distinct characters is: %d\n", 1, lengthOfLongestSubstring("araaci", 1))
	fmt.Printf("The length of the longest substring with %d distinct characters is: %d\n", 3, lengthOfLongestSubstring("cbbebi", 3))
}

func lengthOfLongestSubstring(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}

	maxLen := 0
	left := 0
	currentMap := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		// Add the current character to the window
		currentMap[s[right]]++
		// If we have more than 'k' distinct characters, shrink the window from the left
		for len(currentMap) > k {
			currentMap[s[left]]--
			if currentMap[s[left]] == 0 {
				delete(currentMap, s[left])
				maxLen = currentMap[s[left]]
			}
			left++
		}

		// Update the maximum length of the window
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
