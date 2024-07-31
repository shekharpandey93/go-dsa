package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if arr[i] != arr[0] {
				fmt.Println("from %v", i)
			} else {
				fmt.Println("end %v", i-1)
			}
		}
	}
	if arr[len(arr)-1] != arr[0] {
		fmt.Println("end %v", len(arr)-1)
	}
}
