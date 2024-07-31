package main

import "fmt"

func main()  {
	deleteElement([]int{3,8,12,6,5}, 5)
	idx := largestElement([]int{3,8,12,6,5})
	second := secondLargestElement([]int{3,8,12,6,5})
	isSorted := isArraySorted([]int{1,2,5,4,3})
	fmt.Println("Largest Element Index",idx)
	fmt.Println("Second Largest Element Index",second)
	fmt.Println("is array sorted",isSorted)
	
}

func deleteElement(arr []int, num int)  {
	tempIDX:= -1
	for idx, val := range arr {
		if val == num {
			tempIDX = idx
		}
	}
	if tempIDX != -1 {
		fmt.Println(append(arr[:tempIDX], arr[tempIDX+1:]...))
	}
}

func largestElement(arr []int) int {
	largestIdx := 0
	for idx:=0; idx<len(arr); idx++ {
		if arr[idx] > arr[largestIdx] {
			largestIdx = idx
		}
	}
	return largestIdx
}


func secondLargestElement(arr []int) int {
	resp, largest := -1, 0
	for idx := 1; idx < len(arr); idx++ {
		if arr[idx] > arr[largest] {
			resp = largest
			largest = idx
		} else if arr[idx] != arr[largest] && (resp == -1 || arr[idx] > arr[resp]) {
			resp = idx
		}
	}
	return resp
}

func isArraySorted(arr []int) bool  {
	resp := 0
	for idx := 1; idx < len(arr); idx++ {
		if arr[resp] > arr[idx] {
			return false
		}
		resp = idx
	}
	return true
}
