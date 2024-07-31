package main

import "fmt"

func main() {
	arr := []int{2, 4, 8, 9, 11, 12, 20, 30}
	arr1 := []int{2, 5, 10, 15, 18}
	val := isArrayContainsX(arr, 23)
	fmt.Println("val", val)
	val1 := triplateSum(arr1, 33)
	fmt.Println("val1", val1)
}

func isArrayContainsX(arr []int, x int) bool {
	low := 0
	high := len(arr) - 1
	for low != high {
		if arr[low]+arr[high] == x {
			return true
		} else if arr[low]+arr[high] > x {
			high = high - 1
		} else {
			low = low + 1
		}
	}
	return false
}

func isPair(arr []int, x, idx int) bool {
	low := idx
	high := len(arr) - 1
	for low != high {
		if arr[low]+arr[high] == x {
			return true
		} else if arr[low]+arr[high] > x {
			high = high - 1
		} else {
			low = low + 1
		}
	}
	return false
}

func triplateSum(arr []int, x int) bool {
	for i := 0; i < len(arr)-2; i++ {
		val := isPair(arr, x-arr[0], i+1)
		fmt.Println("return val", val, x-arr[i], i+1)
		if isPair(arr, x-arr[0], i+1) {
			return true
		}
	}
	return false
}
