package main

import "fmt"

func main()  {
	arr := []int{ 1, 2, 3, 5 }
	num := getMissingNumber(arr)
	getMissingNumber2(arr)
	fmt.Println("=======num=====",num)
}

func getMissingNumber(arr []int)  int {
	sum := 1
	for i := 2; i<= len(arr)+ 1; i++ {
		sum += i
		sum -= arr[i-2]
	}
	return sum
}

func getMissingNumber2(arr []int) {
	total := (len(arr) + 1) * (len(arr) + 2) / 2 // sum of all natural number
	fmt.Println("=====total===",total)
	for i := 0; i< len(arr); i++ {
		total -= arr[i] // decrease from array
	}
	fmt.Println(total)

}
