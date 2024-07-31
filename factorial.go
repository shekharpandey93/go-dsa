package main

import "fmt"

func main() {
	fmt.Println(getFactorial(4))
}

func getFactorial(num int) int{
	if num == 0 {
		return 1
	}
	return num * getFactorial(num -1)
}
