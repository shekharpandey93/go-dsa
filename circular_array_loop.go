package main

import (
	"fmt"
)

func circularArrayLoop(nums []int) bool {
	n := len(nums)

	// Helper function to get the next index
	next := func(i int) int {
		return ((i+nums[i])%n + n) % n // Ensures the index is within the bounds of the array
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue // Already visited or considered in another cycle
		}

		slow, fast := i, next(i)
		// Check the loop direction, if it's moving forward or backward consistently
		for nums[fast]*nums[i] > 0 && nums[next(fast)]*nums[i] > 0 {
			if slow == fast {
				if slow == next(slow) {
					break // Single-element loop
				}
				return true // A valid cycle is found
			}
			slow = next(slow)
			fast = next(next(fast))
		}

		// Mark all elements in the current cycle as 0 to avoid revisiting
		slow = i
		val := nums[i]
		for nums[slow]*val > 0 {
			nextIdx := next(slow)
			nums[slow] = 0
			slow = nextIdx
		}
	}

	return false
}

func main() {
	// Test case 1
	nums1 := []int{2, -1, 1, 2, 2}
	fmt.Println(circularArrayLoop(nums1)) // Output: true

	// Test case 2
	nums2 := []int{-1, 2}
	fmt.Println(circularArrayLoop(nums2)) // Output: false

	// Test case 3
	nums3 := []int{-2, 1, -1, -2, -2}
	fmt.Println(circularArrayLoop(nums3)) // Output: false
}
