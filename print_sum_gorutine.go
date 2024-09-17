package main

import "fmt"

func calculateSum(nums []int, result chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	result <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := make(chan int)
	go calculateSum(nums, result)
	fmt.Println("Sum of numbers:", <-result)
}
