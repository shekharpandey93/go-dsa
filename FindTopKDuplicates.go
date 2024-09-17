package main

import (
	"fmt"
	"sort"
)

// CountDuplicates counts occurrences of each integer in the slice
func CountDuplicates(nums []int) map[int]int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	return count
}

// FindTopKDuplicates finds the top k duplicate integers
func FindTopKDuplicates(nums []int, k int) []int {
	// Count occurrences
	count := CountDuplicates(nums)
	// Create a slice of pairs (num, frequency) for duplicates
	var duplicates [][2]int
	for num, freq := range count {
		if freq > 1 {
			duplicates = append(duplicates, [2]int{num, freq})
		}
	}
	// Sort duplicates by frequency (descending) and then by number (ascending)
	sort.Slice(duplicates, func(i, j int) bool {
		if duplicates[i][1] == duplicates[j][1] {
			return duplicates[i][0] < duplicates[j][0]
		}
		return duplicates[i][1] > duplicates[j][1]
	})
	// Get top k duplicates
	result := make([]int, 0, k)
	for i := 0; i < k && i < len(duplicates); i++ {
		result = append(result, duplicates[i][0])
	}

	return result
}

func main() {
	nums := []int{3, 1, 4, 4, 5, 2, 6, 1, 3, 3, 3, 2}
	k := 3

	topK := FindTopKDuplicates(nums, k)
	fmt.Printf("Top %d duplicates: %v\n", k, topK)

	counts := CountDuplicates(nums)
	fmt.Println("Duplicate counts:")
	for num, count := range counts {
		if count > 1 {
			fmt.Printf("%d occurs %d times\n", num, count)
		}
	}
}
