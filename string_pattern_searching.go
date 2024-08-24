package main

import "fmt"

func main() {
	naivePatternSearching("ABCABCD", "ABCD")
}
func naivePatternSearching(txt, pat string) {
	n := len(txt)
	m := len(pat)

	// Slide the pattern over the text one by one
	for i := 0; i <= n-m; {
		j := 0

		// Check if the pattern matches at current position
		for j < m {
			if txt[i+j] != pat[j] {
				break
			}
			j++
		}

		// If the pattern matches, print the current index
		if j == m {
			fmt.Printf("Pattern found at index %d\n", i)
		}

		if j == 0 {
			i++
		} else {
			i = i + j
		}
	}

}
