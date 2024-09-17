package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("Original array:", nums)

	sortColors(nums)
	fmt.Println("Sorted array:", nums)
}

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			mid++
			low++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
