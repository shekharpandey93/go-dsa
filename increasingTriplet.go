package main

import (
	"fmt"
)

func main()  {
	increasingTriplet([]int{0,4,2,1,0,-1,-3})
}

func increasingTriplet(nums []int) bool {
	first := 999999999
	second := 999999999
	for _,n :=range(nums){
		if n <= first{
			first = n
		}else if n <= second{
			second = n
		}else{
			return true
		}
		fmt.Println(first, second)
	}
	return false
}
