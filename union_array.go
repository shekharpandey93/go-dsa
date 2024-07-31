package main

import "fmt"

func main() {
	arr1 := []int{2, 10, 20, 20}
	arr2 := []int{3, 20, 40}
	res := union_array(arr1, arr2)
	fmt.Println(res)
}

func union_array(arr1 []int, arr2 []int) []int {
	idx := 0
	jdx := 0
	result := []int{}
	for idx < len(arr1) && jdx < len(arr2) {
		if idx > 0 && arr1[idx] == arr1[idx-1] {
			idx++
			fmt.Println("111111", idx)
			continue
		}
		if jdx > 0 && arr2[jdx] == arr2[jdx-1] {
			jdx++
			fmt.Println("222222", jdx)
			continue
		}

		if arr1[idx] < arr2[jdx] {
			//result[idx] = arr1[idx]
			result = append(result, arr1[idx])
			fmt.Println("333333", result)
			idx++
			fmt.Println("333333", idx)
		} else if arr1[idx] > arr2[jdx] {
			result = append(result, arr2[jdx])
			jdx++
		} else {
			result = append(result, arr1[idx])
			fmt.Println("5555555", result)
			idx++
			jdx++
		}
	}

	for idx < len(arr1) {
		if idx > 0 && arr1[idx] != arr1[idx-1] {
			result = append(result, arr1[idx])
			fmt.Println("6666666", result)
			idx++
		}
	}
	for jdx < len(arr2) {
		if jdx > 0 && arr2[jdx] != arr2[jdx-1] {
			result = append(result, arr2[jdx])
			fmt.Println("777777", result)
			jdx++
		}
	}
	return result
}
