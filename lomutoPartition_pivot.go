package main

import "fmt"

func main() {
	arr := []int{8, 4, 7, 9, 3, 10, 5}
	//arr := []int{30, 40, 20, 50, 80}
	//reult := partitionQuickSort(arr)
	//fmt.Println(reult)
	quickStort(arr, 0, len(arr)-1)
}

// if pivot is not last then swipe to lasr element and below code will work

func partitionQuickSort(arr []int, left, right int) int {
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
	return idx + 1
}

func quickStort(arr []int, left, right int) {
	if left < right {
		idx := partitionQuickSort(arr, left, right)
		quickStort(arr, left, idx-1)
		quickStort(arr, idx+1, right)
	}
	fmt.Println("arr", arr)
}
