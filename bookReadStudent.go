package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{10, 20, 30, 40}
	res := minPage(arr, len(arr)-1, 2)
	fmt.Println(res)
}

func minPage(arr []int, n, k int) int {
	if k == 1 {
		return sum(arr, 0, n)
	}
	if len(arr) == 1 {
		return arr[0]
	}
	resp := math.MaxInt64
	for i := 1; i < len(arr); i++ {
		resp = min(resp, max(minPage(arr, i, k-1), sum(arr, i, n)))
	}
	return resp
}
func sum(arr []int, b, e int) int {
	sum := 0
	for i := b; i < e; i++ {
		sum += arr[i]
	}
	return sum
}
