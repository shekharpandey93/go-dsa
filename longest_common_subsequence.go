package main

import (
	"fmt"
)

/*
	1.Diff Utility
	2.Minimum insertion and deletion to convert s1 to s2
	3.Shortest common super sequence
	4.Longest Palindrome subsquesnce
	5.Longest repeating subsequence
	6.Space optimized DP solution of LCS
	7. Printing LCS

*/
// Function to find the LCS length using memoization
func lcs(X, Y string, m, n int, memo [][]int) int {
	// If either string is empty, the LCS is 0
	if m == 0 || n == 0 {
		return 0
	}

	// If the value has already been computed, return it
	if memo[m][n] != -1 {
		return memo[m][n]
	}

	// If the last characters match, move both pointers
	if X[m-1] == Y[n-1] {
		memo[m][n] = 1 + lcs(X, Y, m-1, n-1, memo)
	} else {
		// If the last characters don't match, take the maximum of two cases:
		// 1. Exclude the last character of X
		// 2. Exclude the last character of Y
		memo[m][n] = max(lcs(X, Y, m-1, n, memo), lcs(X, Y, m, n-1, memo))
	}

	// Return the memoized value
	return memo[m][n]
}

func main() {
	X := "ABCBDAB"
	Y := "BDCAB"
	m := len(X)
	n := len(Y)

	// Create a memoization table initialized with -1
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// Compute the LCS
	lcsLength := lcs(X, Y, m, n, memo)
	fmt.Printf("The length of the longest common subsequence is %d\n", lcsLength)
}
