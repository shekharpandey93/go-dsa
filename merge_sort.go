package main

import "fmt"

func main() {
	arr := []int{10, 3, 25, 17, 21, 55}
	mergeSort(arr)
}

func mergeSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println(">>>>>>>", arr)
}
