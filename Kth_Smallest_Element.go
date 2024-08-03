package main

import "fmt"

func main() {
	arr := []int{10, 4, 5, 8, 11, 6, 26}
	res := KthSmallest(arr, 5)
	fmt.Println(res)
}
func KthSmallest(arr []int, k int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		p := PartitionQuickSort(arr, low, high)
		fmt.Println("p", p)
		if p == k-1 {
			return arr[p]
		} else if p > k-1 {
			high = p - 1
		} else {
			low = p + 1
		}
	}
	return -1
}

func PartitionQuickSort(arr []int, left, right int) int {
	pivot := arr[right]
	idx := left - 1
	for jdx := left; jdx <= right; jdx++ {
		if arr[jdx] < pivot {
			idx++
			arr[idx], arr[jdx] = arr[jdx], arr[idx]
		}
	}
	if idx == -1 {
		arr[idx+1], arr[right] = arr[right], arr[idx+1]
		return idx + 1
	}
	arr[idx+1], arr[right] = arr[right], arr[idx+1]
	fmt.Println("arr", arr)
	return idx + 1
}
