package main

import "fmt"

func main() {
	arr := []int{2, 4, 1, 3, 5}
	countInversion(arr, 0, len(arr)-1)
}

func countInversion(arr []int, l, r int) int {
	result := 0
	for l < r {
		mid := l + (r-l)/2
		result += countInversion(arr, l, mid)
		result += countInversion(arr, mid+1, r)
		result += countAndMerge(arr, l, mid, r)
	}
	fmt.Println(result)
	return result
}

func countAndMerge(arr []int, start, mid, end int) int {
	fmt.Println(start, mid, end)
	n1 := mid - start + 1
	n2 := end - mid
	newArr := make([]int, n1)
	newArr2 := make([]int, n2)

	for i := 0; i < n1; i++ {
		newArr[i] = arr[start+i]
	}
	for i := 0; i < n2; i++ {
		newArr2[i] = arr[mid+1+i]
	}
	result, left, right, k := 0, 0, 0, 0
	for left < n1 && right < n2 {
		if newArr[left] <= newArr2[right] {
			arr[k] = newArr[left]
			left++
		} else {
			arr[k] = newArr[right]
			right++
			result = result + (n1 - left)
		}
		k++
	}

	for left < n1 {
		arr[k] = newArr[left]
		left++
		k++
	}

	for right < n2 {
		arr[k] = newArr2[right]
		right++
		k++
	}
	fmt.Println(arr)
	return result
}
