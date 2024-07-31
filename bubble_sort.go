package main

import "fmt"

func main() {
	arr := []int {1,4,2,8,5}
	bubbleSort(arr)
}

func bubbleSort(arr []int) {
	swapped := false
	for i:=0; i<len(arr)-1; i++ {
		for j:=0; j<len(arr)-i-1; j++ {
			if arr[j] > arr[j +1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Println("JJJJJJJJ", arr)
}

