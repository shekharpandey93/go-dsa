package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{10, 4, 5, 8, 11, 6, 26}
	res := findKthLargest(arr, 3)
	fmt.Println(res)
}

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		pivotIndex := left + rand.Intn(right-left+1)
		fmt.Println("abc", pivotIndex)
		newPivotIndex := partition(nums, left, right, pivotIndex)
		if newPivotIndex == len(nums)-k {
			return nums[newPivotIndex]
		} else if newPivotIndex > len(nums)-k {
			right = newPivotIndex - 1
		} else {
			left = newPivotIndex + 1
		}
	}
}

func partition(nums []int, left int, right int, pivotIndex int) int {
	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	storedIndex := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[storedIndex] = nums[storedIndex], nums[i]
			storedIndex++
		}
	}
	nums[right], nums[storedIndex] = nums[storedIndex], nums[right]
	return storedIndex
}
