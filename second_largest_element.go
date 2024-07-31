package main

import (
	"fmt"
)

func main()  {
	arr := []int{10, 5, 6, 15}
	index := getSecondLargestElement(arr)
	fmt.Println("ssss", index)
	fmt.Printf("Second Largest Element Is : %v", arr[index])

}
// need to first index is second largest or not.
//
func getSecondLargestElement(arr []int) int  {
	idxOne := 0
	idxTwo :=-1

	for idx := 1; idx < len(arr); idx++ { // loop
		if(arr[idx] > arr[idxOne]) { // comparing index to index
			idxTwo = idxOne // assign first index value to second index
			idxOne = idx // assign index value to first index
			fmt.Println("======", idxOne, idxTwo)
		} else if (idxTwo == -1 || arr[idx] > arr[idxTwo] || arr[idx] != arr[idxOne]) {
			idxTwo = idx
		}
	}
	return idxTwo
}
