package main

import "fmt"

func main() {
	arr := []int{1, 2, 4}
	arr1 := []int{4, 5, 7}

	var freq [101]int
	for i := 0; i < len(arr)+len(arr1); i++ {
		freq[arr[i]]++
		freq[arr1[i+1]]--
	}
	resp := 0
	for i := 1; i < 100; i++ {
		freq[i] = freq[i-1] + freq[i]
		if freq[i] > freq[resp] {
			resp = 1
		}
	}
	fmt.Println(resp)
}
