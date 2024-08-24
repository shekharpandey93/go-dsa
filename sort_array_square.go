package main

import (
	"fmt"
	"math"
)

func main() {
	// Sort array after converting elements to their squares
	arr := []int{-5, -3, -1, 2, 4, 6} // 1,4,9,16,25,36
	squareAndSort(arr)
}

func squareAndSort(arr []int) {
	left := 0
	right := len(arr) - 1
	var temp []int
	//
	for index := len(arr) - 1; index >= 0; index-- {
		if math.Abs(float64(arr[left])) > float64(arr[right]) {
			temp[index] = arr[left] * arr[left]
			left++
		} else {
			temp[index] = arr[right] * arr[right]
			right--
		}
	}

	//var i int
	//// seprate negative and positive element
	//for i =0; i<len(arr); i++ {
	//	if arr[i] > 0 {
	//		break
	//	}
	//}
	//// merge sort logic
	//a := i - 1
	//k := i
	//var temp []int
	//
	//for (a > 0 && k < len(arr)) {
	//	if (arr[a] * arr[a]) < (arr[k] * arr[k]){
	//		temp = append(temp, arr[a] * arr[a])
	//		a--
	//	} else {
	//		temp = append(temp,arr[k] * arr[k])
	//		k++
	//	}
	//	//x++
	//}
	//
	//for a > 0 {
	//	temp = append(temp,arr[a] * arr[a])
	//	a--
	//}
	//
	//for (k < len(arr)) {
	//	temp = append(temp,arr[k] * arr[k])
	//	k++
	//}
	fmt.Println("iiiiii", temp)

}
