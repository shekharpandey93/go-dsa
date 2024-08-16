package main

import "fmt"

func main() {
	str := "ABCDEF"
	str1 := "ADE"
	//output: true
	j := 0
	i := 0
	if len(str) < len(str1) {
		fmt.Printf("%s  is not subsequence of %s", str1, str)
		return
	}

	for i < len(str) && j < len(str1) {
		if str[i] == str1[j] {
			i++
			j++
		} else {
			i++
		}
	}

	if j == len(str1) {
		fmt.Printf("%s  is subsequence of %s", str1, str)
	} else {
		fmt.Printf("%s  is not subsequence of %s", str1, str)
	}
}
