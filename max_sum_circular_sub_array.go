package main

import "fmt"

func main() {
	val := circularMaxSum([]int{8, -4, 3, -5, 4})
	fmt.Println("Max Sum =", val)
}

func getMaxSum(arr []int) int {
	res := arr[0]
	maxSum := arr[0]
	for i := 1; i < len(arr); i++ {
		maxSum = max(arr[i], maxSum+arr[i])
		res = max(res, maxSum)
	}
	return res
}

func circularMaxSum(arr []int) int {
	normalSum := getMaxSum(arr)
	if normalSum < 0 {
		return normalSum
	}
	sunArr := 0
	for i := 0; i < len(arr); i++ {
		sunArr += arr[i]
		arr[i] = -arr[i]
	}
	maxSum := sunArr + getMaxSum(arr)

	return max(normalSum, maxSum)
}
