package main

import "fmt"

func main()  {
	getMaxDiff([]int{2,3,10,6,4,8,1})
}

func getMaxDiff(arr []int)  {
	res := arr[0] - arr[1]
	minval := arr[0]
	for i := 1; i < len(arr); i++ {
		res = max(res, arr[i] - minval)
		minval = min(minval, arr[i])
	}
	fmt.Println("res====",res)
}
