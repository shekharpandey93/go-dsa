package main

import "fmt"

func main()  {
	arr := []int{2,3,5,1,3}
	kidsWithCandies(arr, 3)
}


func kidsWithCandies(candies []int, extraCandies int) []bool {
	var arr []bool;
	var max = 0

	for _, val:=range candies {
		if(max < val) {
			max = val
		}
	}
	fmt.Println("max",max)
	for _, val:=range candies {
		//total := val + extraCandies
		arr = append(arr, val+extraCandies>=max)
		//candies[i]+extraCandies>=max
		//if total >= max {
		//	arr = append(arr, true)
		//} else {
		//	arr = append(arr, false)
		//}
	}
	fmt.Println("val",arr)
	return arr
}
