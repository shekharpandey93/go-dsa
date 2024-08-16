package main

import "fmt"

func main() {
	arr := [][]int{{10, 20, 30}, {15, 25, 35}, {27, 29, 30}}
	searchSorted(arr, 29)
}

func searchSorted(matrix [][]int, target int) bool {
	idx, jdx := 0, len(matrix)-1
	for idx < len(matrix) && jdx >= 0 {
		fmt.Println(idx, jdx)
		if matrix[idx][jdx] == target {
			fmt.Println("FOunt", matrix[idx][jdx])
			return true
		} else if matrix[idx][jdx] > target {
			jdx--
		} else {
			idx++
		}
	}
	fmt.Println("Not FOunt")
	return false
}
