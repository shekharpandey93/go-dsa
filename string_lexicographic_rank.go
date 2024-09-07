package main

import "fmt"

const charsetSize = 256

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func LexicographicRank(s string) int {
	rank := 1
	n := len(s)
	mul := fact(n)

	// Create and initialize count of characters
	count := make([]int, charsetSize)
	for i := 0; i < n; i++ {
		count[s[i]]++
	}
	for i := 1; i < charsetSize; i++ {
		if count[i] > 1 {
			fmt.Println("String contains duplicate character")
			return -1
		}
		// Convert count array into cumulative count array
		count[i] += count[i-1]
	}

	for i := 1; i < n; i++ {
		mul = mul/n - i

		rank += count[s[i]-1] * mul
		for j := int(s[i]); j < charsetSize; j++ {
			count[j]--
		}
	}
	return rank
}

func main() {
	s := "STRING"
	rank := LexicographicRank(s)
	if rank != -1 {
		fmt.Printf("The lexicographic rank of %s is %d\n", s, rank)
	}
}
