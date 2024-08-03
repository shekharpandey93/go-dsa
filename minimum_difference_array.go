package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	arr := []int{10, 8, 1, 4}
	minimumDifference(arr)
}

func minimumDifference(nums []int) int {
	slices.Sort(nums)
	resp := math.MaxInt64
	for i := 1; i < len(nums); i++ {
		resp = min(resp, nums[i]-nums[i-1])
	}
	fmt.Println("resp: ", resp)
	return resp
}
