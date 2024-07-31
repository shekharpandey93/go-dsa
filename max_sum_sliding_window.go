package main

import "fmt"

func main() {
	arr := []int{1, 8, 30, -5, 20, 7}
	k := 4
	curr := 0
	for i := 0; i < k; i++ {
		curr += arr[i]
	}
	resp := curr
	for i := k; i < len(arr); i++ {
		curr += arr[i] - arr[i-k]
		resp = max(resp, curr)
	}
	fmt.Println("resp", resp)
}
