package main

import "fmt"

func main() {
	output := mergeAlternately("abcd", "pq")
	fmt.Println("val", output)
}

func mergeAlternately(word1 string, word2 string) string {

	output := ""
	var remainWord string
	if len(word1) < len(word2) {
		remainWord = word2[len(word1):]
	}
	idx := 0
	for {
		if idx == len(word1) {
			break
		}
		if idx < len(word2) {
			output += string(word1[idx]) + string(word2[idx])
		} else {
			output += string(word1[idx])
		}
		idx++
	}
	return output + remainWord
}
