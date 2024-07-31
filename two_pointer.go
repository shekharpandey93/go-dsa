package main

import "fmt"

func main() {
	arr := []int{2, 5, 8, 12, 30}
	val := twoPointerApproch(arr, 17)
	fmt.Println("val", val)
}

func twoPointerApproch(arr []int, sum int) bool {
	idx, jdx := 0, len(arr)-1

	for idx < jdx {
		if arr[idx]+arr[jdx] == sum {
			return true
		} else if arr[idx]+arr[jdx] > sum {
			jdx--
		} else {
			idx++
		}
	}
	return false
}
