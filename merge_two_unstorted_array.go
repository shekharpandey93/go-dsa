package main

import "fmt"

func main() {
	arr1 := []int{10, 20, 50}
	arr2 := []int{5, 50, 50}
	resp := mergeTwoArrays(arr1, arr2)
	fmt.Println(resp)

}

func mergeTwoArrays(arr1 []int, arr2 []int) []int {
	result := []int{}
	idx, jdx := 0, 0
	for idx < len(arr1) && jdx < len(arr2) {
		if arr1[idx] <= arr2[jdx] {
			//result[idx] = arr1[idx]
			result = append(result, arr1[idx])
			//fmt.Println(arr1[idx])
			idx++
		} else {
			//result[jdx] = arr2[jdx]
			result = append(result, arr2[jdx])
			//fmt.Println(arr2[jdx])
			jdx++
		}
	}
	for idx < len(arr1) {
		//fmt.Println("result1", result)
		result = append(result, arr1[idx])
		//fmt.Println(arr1[idx])
		idx++
	}
	for jdx < len(arr2) {
		//fmt.Println("result2", result)
		result = append(result, arr2[jdx])
		//fmt.Println(arr2[jdx])
		jdx++
	}
	//fmt.Println("result3", result)
	return result
}
