package main

import "fmt"

func main() {
	arr := []int{8, 58, 71, 18, 31, 32, 63, 92, 43, 3, 91, 93, 25, 80, 28}
	var arr1 []int
	idx := 0
	fmt.Println("=====1")
	var idx2 = len(arr)
	tempIdx := idx + 1
	fmt.Println("leng", idx2)
	for {
		fmt.Println("=====2", idx)
		fmt.Println("arr[idx]", tempIdx)
		val := 0
		if tempIdx == idx2 {
			tempIdx = idx + 1
		}
		if arr[idx] < arr[tempIdx]{
			fmt.Println("=====3")
			val = getLestVal(arr, arr[idx], arr[tempIdx], idx)
			tempIdx = idx + 1
			fmt.Println("=====4", val)
			idx++
		} else {
			tempIdx++
			fmt.Println("tempIdx", tempIdx)
		}
		arr1 = append(arr1, val)
		if idx > idx2 {
			break
		}
	}
	fmt.Println("=====", arr1)
}

func getLestVal(arr []int, val, val1, idx int) int {
	if val1 > arr[idx+1] && val < arr[idx+1] {
		fmt.Println("=====", val1)
		return getLestVal(arr, val, arr[idx+1], idx+1)
	}
	return val1
}
