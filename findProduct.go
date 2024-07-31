package main

import "fmt"

func findProduct(nums []int) []int {
	n := len(nums)

	// Initialize prefix and suffix arrays
	prefix := make([]int, n)
	suffix := make([]int, n)

	// Calculate prefix products
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	fmt.Println("prefix",prefix)
	// Calculate suffix products
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	fmt.Println("suffix",suffix)
	// Calculate the final product array
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := findProduct(nums)
	fmt.Println("Original array:", nums)
	fmt.Println("Product array:", result)
}
