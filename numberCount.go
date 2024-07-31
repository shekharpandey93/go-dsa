package main

import "fmt"

func main()  {
	count := numberCount(2123457)
	fmt.Println("count is",count)
}

func numberCount(num int)  int{
	count := 0

	for num%10 != 0 {
		num = num/10
		count++
	}
	return count
}
