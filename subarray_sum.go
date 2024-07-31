package main

import "fmt"

//N = 10, S = 15
//A[] = {1,2,3,4,5,6,7,8,9,10}

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	getSubArraySub(arr, 15, 10)
}

func getSubArraySub(arr []int, sum, n int)  {
	currSum := sum
	//ar subArr []int
	start :=0
	for idx:=1; idx< n; idx++ {
		if idx < n {
			currSum = currSum + arr[idx]
		}
		for currSum > sum && start < idx -1 {
			currSum = currSum - arr[start]
			start++

			fmt.Println("=currSum====",currSum)
			fmt.Println("=start====",start)
		}

		if currSum == sum {

		}

	}
}
