package main

import "fmt"

func main() {
	arr := []int{3, 4, 8, -9, 9, 7}
	fmt.Println(isEquilibriumPoint(arr))
}

func isEquilibriumPoint(arr []int) bool {
	ls := 0
	rs := 0
	for i := 0; i < len(arr); i++ {
		rs += arr[i]
	}

	for i := 0; i < len(arr); i++ {
		rs -= arr[i]
		if ls == rs {
			return true
		}
		ls += arr[i]
	}
	return false
}
