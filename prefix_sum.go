package main

import "fmt"

func main() {
	arr := []int{2, 8, 3, 9, 6, 5, 4}
	sumArr := make([]int, len(arr))
	sumArr[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		sumArr[i] = sumArr[i-1] + arr[i]
	}
	sumWArr := make([]int, len(arr))
	sumWArr[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		sumWArr[i] = (i + 1) * (sumWArr[i-1] + arr[i])
	}
	fmt.Println("sumWArr", sumWArr)
	sum := getSum(sumArr, 0, 2)
	fmt.Println("sum:", sum)
	sum = getSum(sumArr, 2, 6)
	fmt.Println("sum1:", sum)
}

func getSum(arr []int, l, r int) int {
	if l == 0 {
		return arr[r]
	}
	return arr[r] - arr[l-1]
}

func getWSum(arr []int, l, r int) int {
	if l == 0 {
		return arr[r]
	}
	return arr[r] - arr[l-1]
}
