package main

import "fmt"

func main() {
	arr := []int{5, 20, 40, 30, 20, 50, 60}
	val := getPickElement(arr)
	fmt.Println(arr[val])
}

func getPickElement(arr []int) int {
	low := 0
	heigh := len(arr) - 1
	for low <= heigh {
		mid := (low + heigh) / 2
		fmt.Println("mid", mid)
		if (mid == 0 || arr[mid] >= arr[mid-1]) && (mid == len(arr)-1 || arr[mid] >= arr[mid+1]) {
			return mid
		} else if mid > 0 && arr[mid-1] >= arr[mid] {
			heigh = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
