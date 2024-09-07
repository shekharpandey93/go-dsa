package main

import "fmt"

func main() {
	arr := []int{10, 4, 5, 8, 11, 6, 26}
	res := findKthLargest(arr, 3)
	fmt.Println(res)
}
func findKthLargest(nums []int, k int) int {
	left, right, k := 0, len(nums)-1, len(nums)-k

	for left <= right {
		boundary1, boundary2 := getPivots(nums, left, right)

		if k >= boundary1 && k <= boundary2 {
			return nums[k]
		} else if boundary2 < k {
			left = boundary2 + 1
		} else {
			right = boundary1 - 1
		}
	}
	return -1
}

func getPivots(nums []int, left int, right int) (int, int) {
	mid := left
	pivot := nums[left]

	for mid <= right {
		if nums[mid] < pivot {
			nums[mid], nums[left] = nums[left], nums[mid]
			left++
			mid++
		} else if nums[mid] == pivot {
			mid++
		} else {
			nums[right], nums[mid] = nums[mid], nums[right]
			right--
		}
	}

	return left, mid - 1
}
