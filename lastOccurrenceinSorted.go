package main

import "fmt"

func main() {
	//[10,15,20,20,40,40]
	//x = 20
	val := lastOccurrenceInSorted([]int{10, 15, 20, 20, 40, 40}, 20)
	fmt.Println(val)
}

func lastOccurrenceInSorted(arr []int, x int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			if (mid != len(arr)-1) && arr[mid] != arr[mid+1] {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
