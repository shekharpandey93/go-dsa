package main

import "fmt"

func main()  {
	arr := moveAllZeroEnd([]int{10,8,0,0,12,0})
	fmt.Println("arr",arr)
}


func moveAllZeroEnd(arr []int) []int{
	count := 0
	for idx := 0; idx < len(arr); idx++ {
		if arr[idx] != 0 {
			arr[idx], arr[count] =  arr[count], arr[idx]
			count++
		}
	}
	return arr
}
