package main

import "fmt"

func main() {
	numSubset([]int{1,2,3}, []int{},0, 4, 0)
	//val := numSubsetNew([]int{1,2,3},len([]int{1,2,3}),4)
	//fmt.Println("Radhe Radhe", val)
}
//func numSubsetNew(arr []int, idx, sum int) int {
//	//var resp []string
//	if idx == 0{
//		returnVal := 0
//		if sum != 0 {
//			returnVal = 1
//		}
//		return returnVal
//	}
//	return numSubsetNew(arr,idx - 1, sum) + numSubsetNew(arr, idx-1, sum - arr[idx-1])
//}

func numSubset(arr []int, current []int, idx, sum, count int) {
	//var resp []string
	if idx == len(arr){
			if sumCount(current, sum) {
				count++
			}
			fmt.Println("count",count)
		return
	}
	numSubset(arr,current, idx+1, sum, count)
	current = append(current, arr[idx])
	numSubset(arr,current, idx+1, sum, count)
}

func sumCount(arr []int, sum int) bool {
	tem := 0
	for idx:= 0; idx <= len(arr)-1; idx++ {
		tem += arr[idx]
	}
	return tem == sum
}
