package main

import "fmt"

func main()  {
	allPermutation("abc", 0)
}

func allPermutation(str string, idx int)  {
	if idx == len(str) - 1 {
		fmt.Println(str)
		return
	}
	bytes := []byte(str)
	for jdx := 0; jdx <= len(str) - 1; jdx++ {
		bytes[idx], bytes[jdx] = bytes[jdx], bytes[idx]
		allPermutation(string(bytes), idx+1)
		bytes[idx], bytes[jdx] = bytes[jdx], bytes[idx]
	}
}


