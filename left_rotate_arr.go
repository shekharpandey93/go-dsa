package main

import "fmt"

func main()  {
	arr := []int{1,2,3,4,5}
	d := 3
	arr = reversArr(arr, 0, d- 1)
	arr = reversArr(arr, d, len(arr) - 1)
	arr = reversArr(arr, 0,  len(arr) - 1)
	fmt.Println("arr",arr)

}

func reversArr(arr []int, low, high int)  []int{
	for low < high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
	return arr
}
