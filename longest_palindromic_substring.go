package main

import "fmt"

func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// return the palindrome substring
	return s[left+1 : right]
}
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	longest := ""

	for i := 0; i < len(s); i++ {
		oddPalindrome := expandAroundCenter(s, i, i)
		evenPalindrome := expandAroundCenter(s, i, i+1)
		if len(oddPalindrome) > len(longest) {
			longest = oddPalindrome
		}
		if len(evenPalindrome) > len(longest) {
			longest = evenPalindrome
		}
	}
	return longest
}

func main() {
	s := "babad"
	fmt.Println("Longest Palindromic Substring:", longestPalindrome(s)) // Output: "bab" or "aba"

	s2 := "cbbd"
	fmt.Println("Longest Palindromic Substring:", longestPalindrome(s2)) // Output: "bb"
}
