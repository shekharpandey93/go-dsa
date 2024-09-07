package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWordOccurrences(paragraph string) map[string]int {
	// Regular expression to match non-alphabetic characters
	re := regexp.MustCompile(`[^\w\s]`)
	// Replace all non-alphabetic characters with an empty string
	cleanedParagraph := re.ReplaceAllString(paragraph, "")

	// Convert to lowercase to ensure case insensitivity
	cleanedParagraph = strings.ToLower(cleanedParagraph)

	// Split the paragraph into words
	words := strings.Fields(cleanedParagraph)

	// Create a map to count the occurrences of each word
	wordCount := make(map[string]int)

	// Count the occurrences
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	paragraph := `
		In computer science, a data structure is a data organization, management, 
		and storage format that enables efficient access and modification. More 
		precisely, a data structure is a collection of data values, the relationships 
		among them, and the functions or operations that can be applied to the data.
	`
	wordCount := countWordOccurrences(paragraph)

	// Print the word count
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
